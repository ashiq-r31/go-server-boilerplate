package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "You can now get started on your app!"}`))
	})

	fmt.Println("Your app is now serving at 127.0.0.1:3000")

	http.ListenAndServe(":3000", r)
}