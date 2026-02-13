package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Println("Servidor iniciado en localhost:8070...")
	http.ListenAndServe(":8070", nil)
}
