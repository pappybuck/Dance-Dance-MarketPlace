using Janus.Identity;
using Janus.Models;
using Janus.Services;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.ModelBinding;
using System.Data.Entity;
using System.Net;

namespace Janus.Controllers
{
    [Route("auth/")]
    [ApiController]
    public class AuthController : ControllerBase
    {
        private readonly SignInManager<User> _signInManager;
        private readonly UserManager<User> _userManager;
        private readonly RoleManager<IdentityRole> _roleManager;
        private readonly DbContext _dbcontext;


        public AuthController(UserManager<User> userManager, SignInManager<User> signInManager, RoleManager<IdentityRole> roleManager, DbContext dbContext)
        {
            _userManager = userManager;
            _signInManager = signInManager;
            _roleManager = roleManager;
            _dbcontext = dbContext;
        }

        [HttpPost("register")]
        public async Task<IActionResult> Register([FromBody] UserRegistrationModel request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest("User Registration Failed");
            }
            var user = new User()
            {
                Email = request.Email,
                FirstName = request.FirstName,
                LastName = request.LastName,
                UserName = request.UserName,
            };
            var result = await _userManager.CreateAsync(user, request.Password);
            if (!result.Succeeded)
            {
                var dictionary = new ModelStateDictionary();
                foreach (IdentityError error in result.Errors)
                {
                    dictionary.AddModelError(error.Code, error.Description);
                }

                return new BadRequestObjectResult(new { Message = "User Registration Failed", Errors = dictionary });
            }
            _dbcontext.Profiles.Add(new Profile
            {
                Email = request.Email,
                FirstName = request.FirstName,
                LastName = request.LastName,
                User = user,
                Phone = "N/A"
            });
            _dbcontext.SaveChanges();
            return Ok("User Registration Successful");
        }
        [HttpPost("register-admin")]
        public async Task<IActionResult> RegisterAdmin([FromBody] UserRegistrationModel request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest("User Registration Failed");
            }
            var user = new User()
            {
                Email = request.Email,
                FirstName = request.FirstName,
                LastName = request.LastName,
                UserName = request.UserName,
            };
            var result = await _userManager.CreateAsync(user, request.Password);
            if (!result.Succeeded)
            {
                var dictionary = new ModelStateDictionary();
                foreach (IdentityError error in result.Errors)
                {
                    dictionary.AddModelError(error.Code, error.Description);
                }

                return new BadRequestObjectResult(new { Message = "User Registration Failed", Errors = dictionary });
            }
            await _userManager.AddToRoleAsync(user, "Admin");
            _dbcontext.Profiles.Add(new Profile
            {
                Email = request.Email,
                FirstName = request.FirstName,
                LastName = request.LastName,
                User = user,
                Phone = "N/A"
            });
            _dbcontext.SaveChanges();
            return Ok("User Registration Successful");
        }

        [HttpPost("login")]
        public async Task<IActionResult> Login([FromBody] UserLoginModel request)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest("Unable to login");
            }
            var user = await _userManager.FindByEmailAsync(request.Email);
            if (user == null)
            {
                return Unauthorized("Login failed");
            }
            var result = await _signInManager.CheckPasswordSignInAsync(user, request.Password, false);
            if (result.Succeeded)
            {
                var userRoles = await _userManager.GetRolesAsync(user);
                var authClaims = authenicationClaims(user);
                HttpClient client = new HttpClient();
                var url = Environment.GetEnvironmentVariable("JWT_URL") ?? "http://localhost:8080";
                var response = await client.PostAsJsonAsync(url + "/jwt",
                    new TokenRequestModel
                    {
                        email = user.Email,
                        roles = authClaims,
                    });
                var content = response.Content.ReadFromJsonAsync<TokenResponseModel>();
                if (content.Result != null)
                {
                    var token = content.Result.token;
                    SetAccessToken(token);
                    var refreshToken = TokenService.GenerateRefreshToken(content.Result.id);
                    if (user.token == null)
                    {
                        user.token = new List<RefreshToken?>{ refreshToken };
                    }
                    else
                    {
                        user.token.Add(refreshToken);
                    }
                    SetRefreshToken(refreshToken.Token);
                    await _dbcontext.SaveChangesAsync();
                    return Ok( new ReturnModel
                    {
                        email = user.Email,
                        firstName = user.FirstName,
                        lastName = user.LastName,
                    });
                }
                return BadRequest("Unable to login");
            }
            
            return Unauthorized("Login failed");
        }

        [HttpGet("login")]
        public async Task<IActionResult> Login()
        {
            HttpClient client = new HttpClient();
            var auth = Request.Cookies["Authorization"];
            if (auth == null || auth.Split(" ").Length != 2)
            {
                return Unauthorized("Invalid Token");
            }
            client.DefaultRequestHeaders.Add("Authorization", auth);
            var url = Environment.GetEnvironmentVariable("JWT_URL") ?? "http://localhost:8080";
            var response = await client.GetAsync(url + "/verify");
            if (response.StatusCode == HttpStatusCode.OK 
                || response.StatusCode == HttpStatusCode.NotAcceptable
                || response.StatusCode == HttpStatusCode.Accepted)
            {
                var refresh = Request.Cookies["Refresh"];
                if (refresh != null)
                {
                    var user = _dbcontext.Users.Include(u => u.token).FirstOrDefault(u => u.token.Any(t => t.Token == refresh));
                    if (user != null)
                    {
                        var id = response.Content.ReadAsStringAsync().Result;
                        id = id.Replace("\"", "");
                        var refreshToken = _dbcontext.RefreshToken.FirstOrDefault(t => t.Token == refresh);
                        if (refreshToken != null && refreshToken.id == id 
                                         && refreshToken.TokenExpire > DateTime.UtcNow 
                                         && refreshToken.Invalidated == false)
                        {
                                refreshToken.Invalidated = true;
                                refreshToken.Revoked = DateTime.UtcNow;
                                var userRoles = await _userManager.GetRolesAsync(user);
                                var authClaims = authenicationClaims(user);
                                url = Environment.GetEnvironmentVariable("JWT_URL") ?? "http://localhost:8080";
                                response = await client.PostAsJsonAsync(url + "/jwt",
                                    new TokenRequestModel
                                    {
                                        email = user.Email,
                                        roles = authClaims,
                                    });
                                var content = response.Content.ReadFromJsonAsync<TokenResponseModel>();
                                if (content.Result != null)
                                {
                                    SetAccessToken(content.Result.token);
                                    var token = content.Result.token;
                                    refreshToken = TokenService.GenerateRefreshToken(content.Result.id);
                                    user.token.Add(refreshToken);
                                    SetRefreshToken(refreshToken.Token);
                                    await _dbcontext.SaveChangesAsync();
                                    return Ok( new ReturnModel
                                    {
                                        email = user.Email,
                                        firstName = user.FirstName,
                                        lastName = user.LastName,
                                    });
                                }
                        }
                    }

                    return BadRequest();
                }   
            }
            return Unauthorized();
            
        }

        [HttpGet("Verify")]
        public async Task<IActionResult> Verify()
        {
            HttpClient client = new HttpClient();
            var auth = Request.Cookies["Authorization"];
            if (auth == null || auth.Split(" ").Length != 2)
            {
                return Unauthorized("Invalid Token");
            }
            client.DefaultRequestHeaders.Add("Authorization", auth);
            var queryString = Request.QueryString;
            var url = Environment.GetEnvironmentVariable("JWT_URL") ?? "http://localhost:8080";
            var response = await client.GetAsync(url + "/verify" + queryString);
            if (response.StatusCode == HttpStatusCode.OK)
            {
                return Ok();
            }
            if (response.StatusCode == HttpStatusCode.NotAcceptable || response.StatusCode == HttpStatusCode.Accepted)
            {
                var refresh = Request.Cookies["Refresh"];
                if (refresh != null)
                {
                    var user = _dbcontext.Users.Include(u => u.token).FirstOrDefault(u => u.token.Any(t => t.Token == refresh));
                    if (user != null)
                    {
                        var id = response.Content.ReadAsStringAsync().Result;
                        id = id.Replace("\"", "");
                        var refreshToken = _dbcontext.RefreshToken.FirstOrDefault(t => t.Token == refresh);
                        if (refreshToken != null && refreshToken.id == id 
                                                 && refreshToken.TokenExpire > DateTime.UtcNow 
                                                 && refreshToken.Invalidated == false)
                        {
                            refreshToken.Invalidated = true;
                            refreshToken.Revoked = DateTime.UtcNow;
                            var userRoles = await _userManager.GetRolesAsync(user);
                            var authClaims = authenicationClaims(user);
                            response = await client.PostAsJsonAsync(url + "/jwt" ,
                                new TokenRequestModel
                                {
                                    email = user.Email,
                                    roles = authClaims,
                                });
                            var content = response.Content.ReadFromJsonAsync<TokenResponseModel>();
                            if (content.Result != null)
                            {
                                var token = content.Result.token;
                                SetAccessToken(token);
                                refreshToken = TokenService.GenerateRefreshToken(content.Result.id);
                                user.token.Add(refreshToken);
                                SetRefreshToken(refreshToken.Token);
                                await _dbcontext.SaveChangesAsync();
                                return Ok();
                            }
                        }
                    }
                
                    return BadRequest();
                }   
            }
            
            return Unauthorized();
           
        }
        
        [HttpGet("logout")]
        public async Task<IActionResult> Logout()
        {
            var refresh = Request.Cookies["Refresh"];
            Response.Cookies.Delete("Authorization", new CookieOptions{ Expires = DateTime.Now.AddDays(-1) });
            Response.Cookies.Delete("Refresh", new CookieOptions { Expires = DateTime.Now.AddDays(-1) });
            if (!string.IsNullOrEmpty(refresh))
            {
                var user = _dbcontext.Users.Include(u => u.token).FirstOrDefault(u => u.token.Any(t => t.Token == refresh));
                if (user != null)
                {
                    var token = _dbcontext.RefreshToken.FirstOrDefault(t => t.Token == refresh);
                    if (token != null)
                    {
                        token.Invalidated = true;
                        token.Revoked = DateTime.UtcNow;
                    }
                    _dbcontext.SaveChanges();
                    return Ok();
                }
            }
            return BadRequest();
        }

        private void SetAccessToken(string token)
        {
            Response.Cookies.Append("Authorization", $"Bearer {token}", new CookieOptions
            {
                HttpOnly = true,
                Expires = DateTime.UtcNow.AddDays(7),
                SameSite = SameSiteMode.None,
                Secure = true,
                Domain = "patrickbuck.net"
            });
        }

        private void SetRefreshToken(string token)
        {
            Response.Cookies.Append("Refresh", token, new CookieOptions
            {
                HttpOnly = true,
                Expires = DateTime.UtcNow.AddDays(7),
                SameSite = SameSiteMode.None,
                Secure = true,
                Domain = "patrickbuck.net"
            });
        }

        private String[] authenicationClaims(User user)
        {
            var userRoles = _userManager.GetRolesAsync(user).Result;
            string[] authClaims = new string[userRoles.Count];
            for (int i = 0; i < userRoles.Count; i++)
            {
                authClaims[i] = userRoles[i];
            }
            return authClaims;
        }
    }
}
