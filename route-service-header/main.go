package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

const (
	DEFAULT_PORT     = "8080"
	CF_FORWARDED_URL = "X-Cf-Forwarded-Url"
	CF_CANARY_URL    = "X-Cf-Canary-Url"
)

func main() {
	log.SetOutput(os.Stdout)
	http.Handle("/", newProxy())
	log.Fatal(http.ListenAndServe(":"+getPort(), nil))
}

func newProxy() http.Handler {
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {

			canaryURL := req.Header.Get(CF_CANARY_URL)
			if canaryURL != "" {
				url, err := url.Parse(canaryURL)
				if err != nil {
					log.Fatalln(err.Error())
				}
				req.URL = url
				req.Host = url.Host

			} else {
				forwardedURL := req.Header.Get(CF_FORWARDED_URL)

				url, err := url.Parse(forwardedURL)
				if err != nil {
					log.Fatalln(err.Error())
				}

				req.URL = url
				req.Host = url.Host
			}

		},
	}
	return proxy
}

func getPort() string {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
	return port
}

func getEnv(env string, defaultValue int) int {
	var (
		v      string
		config int
	)
	if v = os.Getenv(env); len(v) == 0 {
		return defaultValue
	}

	config, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return config
}
