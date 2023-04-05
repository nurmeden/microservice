package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	service1URL, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	service2URL, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	service1Proxy := httputil.NewSingleHostReverseProxy(service1URL)
	service2Proxy := httputil.NewSingleHostReverseProxy(service2URL)

	http.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		service1Proxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		service2Proxy.ServeHTTP(w, r)
	})

	log.Println("API Gateway is listening on :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
