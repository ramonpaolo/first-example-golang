package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Println("Hello, World!")
	})

	http.ListenAndServe(":3000", nil)
}
