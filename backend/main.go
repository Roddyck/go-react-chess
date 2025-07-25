package main

import (
	"log"
	"net/http"

	"github.com/Roddyck/go-react-chess/backend/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.AllowCors,
	)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: stack(mux),
	}

	log.Fatal(httpServer.ListenAndServe())
}
