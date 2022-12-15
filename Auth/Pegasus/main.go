package main

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwk"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

/*
A struct containing the additional claims that are added to a JWT
*/
type Claims struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

/*
A struct that represents the body of the request needed to create a JWT
*/
type JwtRequest struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

/*
Initializes the Public and Private RSA keys
*/
func init() {
	signBytes, err := ioutil.ReadFile("./RSA/id_rsa")
	if err != nil {
		log.Fatal(err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}
	verifyBytes, err := ioutil.ReadFile("./RSA/id_rsa.pub")
	if err != nil {
		log.Fatal(err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())

	// Creates a POST route that creates JWTs
	r.POST("/jwt", jwtHandler)
	// Creates a GET route that verifies JWTs and notes if they need to be refreshed
	r.GET("/verify", verifyHandler)
	// Creates a GET route for the Json Web Key created using the Public Key
	r.GET("/jwks", jwksHandler)

	r.Run()

}

/*
A handler that is used to Create Jwts given a JwtRequest struct as the body of the Request.
Returns the jwt and the id of the token.
*/
func jwtHandler(c *gin.Context) {
	var jwtRequest JwtRequest
	err := c.BindJSON(&jwtRequest)
	if err != nil {
		c.JSON(400, "Bad Request")
		return
	}
	id := uuid.New().String() // This value is stored with refresh token inside of the database.
	claims := Claims{
		Email: jwtRequest.Email,
		Roles: jwtRequest.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Patrickbuck.net",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * 60 * time.Second)), //Each JWT is only valid for 5 minutes after being issued
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"token": signedToken,
		"id":    id,
	})
}

/*
A handler that is used to verify the validity and expiration of tokens.
The token is stored in the Authorization header and begins with 'Bearer '.
The verify route can contain a roles query which contains valid roles that are to used to check the token.
Return status codes: 401: Unauthorized, 400: Bad Request, 403: Invalid token, 406: Expiered but valid token,
202: Valid but about to expire token, 200: Valid token.
*/
func verifyHandler(c *gin.Context) {
	expired := false
	header := c.Request.Header["Authorization"]
	if len(header) == 0 {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	tokenString := strings.Split(header[0], " ")
	if len(tokenString) != 2 {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	token, err := jwt.ParseWithClaims(tokenString[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		if !strings.Contains(err.Error(), "token is expired") {
			//If the only error is that the token is expired, then continue with validation request.
			c.JSON(403, gin.H{
				"Messsage": "Forbidden",
			})
			return
		}
		expired = true
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	roles := c.QueryArray("roles")
	if roles == nil {
		if expired {
			c.JSON(406, token.Claims.(*Claims).ID)
			return
		}
		if time.Now().Add(2 * time.Minute).After(claims.ExpiresAt.Time) {
			c.JSON(202, token.Claims.(*Claims).ID)
			return
		}
		c.JSON(200, token.Claims.(*Claims).ID)
		return
	}
	for _, role := range roles {
		if contains(claims.Roles, role) {
			if expired {
				c.JSON(406, token.Claims.(*Claims).ID)
				return
			}
			if time.Now().Add(2 * time.Minute).After(claims.ExpiresAt.Time) {
				c.JSON(202, token.Claims.(*Claims).ID)
				return
			}
			c.JSON(200, token.Claims.(*Claims).ID)
			return
		}
	}
	c.JSON(401, "Unauthorized")

}

/*
Helper function to check if a string is inside of an array.
*/
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/*
A handler that outputs a Json Web Key set in proper format.
*/
func jwksHandler(c *gin.Context) {
	key, err := jwk.New(verifyKey)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	}

	key.Set(jwk.KeyUsageKey, "sig")
	key.Set(jwk.AlgorithmKey, "RS256")
	key.Set(jwk.KeyIDKey, "VerifyKey")

	jwks := jwk.NewSet()
	jwks.Add(key)

	c.JSON(200, jwks)
}
