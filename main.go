package main

import (
	"bytes"
	"net/http"
)

func handle(rw http.ResponseWriter, req *http.Request) {
	if _, err := bytes.NewBufferString("Hello, World!").WriteTo(rw); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
