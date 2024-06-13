package main

import (
	"fmt"
	"net/http"
)

const PORT = 8080

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))

	s := &http.Server{
		Addr:    fmt.Sprintf("%s%d", ":", PORT),
		Handler: mux,
	}
	s.ListenAndServe()
	fmt.Println("Server started")
}
