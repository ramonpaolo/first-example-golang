package main

import (
	"fmt"
	"net/http"
	requests "simple-server/src"
)

func main() {
	fmt.Println("Iniciando...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var n string

		fmt.Println("Hello, World!")
		fmt.Println(r.URL.Path)

		n, err := requests.GetUser("")

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))

			return
		}

		w.WriteHeader(200)
		w.Write([]byte(n))
	})

	http.ListenAndServe(":3000", nil)
}
