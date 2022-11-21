package test

import (
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	http.Handle("/hello1", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world1"))
	}))
	http.Handle("/hello2", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world2"))
	}))
	http.ListenAndServe("localhost:8088", nil)
}
