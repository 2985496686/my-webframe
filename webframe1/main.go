package main

import (
	"log"
	"webframe1/gee"
)

func main() {
	r := gee.New()
	//r.Get("/hello")
	log.Fatal(r.Run(""))
}
