package main

import (
	"log"

	"github.com/Roddyck/go-react-chess/backend/internal/api"
)

func main() {
	log.Fatal(api.StartServer("8080"))
}

