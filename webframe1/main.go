package main

import (
	"fmt"
	"log"
	"net/http"
	"webframe1/gee"
)

func main() {
	gee := gee.New()
	gee.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world --- get")
	})

	gee.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello --- post")
	})
	log.Fatal(gee.Run("localhost:9998"))
}
