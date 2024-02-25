package main

import (
	"diplom/api"
	"log"
	"net/http"
)

func main() {
	log.Println("Запуск сервера")
	r := api.CreateRouter()
	http.ListenAndServe("localhost:8080", r)
	/*
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Printf("Can`t run server:\n%v", err)
			os.Exit(1)
		}

	*/
}
