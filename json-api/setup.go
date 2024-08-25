package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello World"))
	fmt.Fprintf(w, "Hello World")
}

func Setup() {
	http.HandleFunc("GET /hello", sayHello)
}
