package main

import (
	"diplom/api/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Start server")
	r := routes.CreateRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Can`t run server:\n%v", err)
		os.Exit(1)
	}
}
