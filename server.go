package main

import (
	"fmt"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
	Hour    int
}

func main() {

	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/", fs)

	fmt.Println("Server is listenng...")
	http.ListenAndServe(":8181", nil)
}
