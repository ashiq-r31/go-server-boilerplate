package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "You can now get started on your app!"}`))
}

func main() {
	r := chi.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./static")))
	r.Get("/example", handler)

	fmt.Println("Your app is now serving at 127.0.0.1:3000")

	http.ListenAndServe(":3000", r)
}