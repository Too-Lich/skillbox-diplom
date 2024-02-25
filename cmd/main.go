package main

import (
	"diplom/api/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Запуск сервера")
	r := routes.CreateRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Can`t run server:\n%v", err)
		os.Exit(1)
	}
}
