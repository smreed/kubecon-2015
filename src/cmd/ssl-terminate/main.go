package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	upstream, err := url.Parse("http://localhost")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(upstream)
	http.Handle("/", proxy)
	log.Fatal(http.ListenAndServeTLS(":443", "/cert.pem", "/key.pem", nil))
}
