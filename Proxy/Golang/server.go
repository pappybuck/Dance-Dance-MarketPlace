package main

import (
	"io/ioutil"
	"log"

	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Concurrent int      `yaml:"Concurrent"`
	Servers    []Server `yaml:"Servers"`
}

type Server struct {
	Route          string `yaml:"route"`
	Host           string `yaml:"host"`
	Authentication string `yaml:"authentication"`
}

func main() {

	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Error reading config file")
	}
	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	if len(config.Servers) == 0 {
		log.Fatal("No servers found in config")
	}
	if config.Concurrent <= 0 {
		config.Concurrent = 1000
	}

	s := fasthttp.Server{
		Concurrency: config.Concurrent,
	}

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		for _, server := range config.Servers {
			if string(ctx.Host()) == server.Route {
				var auth []byte = nil
				var refresh []byte = nil
				if server.Authentication != "" {
					req := fasthttp.AcquireRequest()
					ctx.Request.CopyTo(req)
					req.SetRequestURI(server.Authentication)
					req.Header.SetMethod("GET")
					req.Header.SetCookieBytesKV([]byte("Authorization"), ctx.Request.Header.Cookie("Authorization"))
					req.Header.SetCookieBytesKV([]byte("Refresh"), ctx.Request.Header.Cookie("Refresh"))
					err := fasthttp.Do(req, &ctx.Response)
					if err != nil {
						log.Default().Println(err)
						s.Shutdown()
					}
					if ctx.Response.StatusCode() == 401 {
						ctx.Redirect("https://www.google.com", 302)
						return
					}
					auth = ctx.Response.Header.PeekCookie("Authorization")
					refresh = ctx.Response.Header.PeekCookie("Refresh")
				}
				ctx.Request.SetHost(server.Host)
				err := fasthttp.Do(&ctx.Request, &ctx.Response)
				if err != nil {
					log.Default().Println(err)
					ctx.SetStatusCode(500)
					ctx.SetBody([]byte("Internal Server Error"))
					return
				}
				if auth != nil {
					cookie1 := fasthttp.AcquireCookie()
					cookie1.SetKey("Authorization")
					cookie1.SetValueBytes(auth)
					ctx.Response.Header.SetCookie(cookie1)
					cookie2 := fasthttp.AcquireCookie()
					cookie2.SetKey("Refresh")
					cookie2.SetValueBytes(refresh)
					ctx.Response.Header.SetCookie(cookie2)
				}
				return
			}
		}
		ctx.SetStatusCode(404)
		ctx.SetBody([]byte("Not Found"))
	}

	s.Handler = requestHandler

	log.Default().Println("Starting server on :4000")

	if err := s.ListenAndServe(":4000"); err != nil {
		log.Default().Println(err)
		s.Shutdown()
	}
}
