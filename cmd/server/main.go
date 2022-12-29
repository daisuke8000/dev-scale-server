package main

import (
	"dev-scale-server/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
