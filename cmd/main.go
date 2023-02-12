package main

import (
	"github.com/CyclopsV/cities-informer-skillbox/api"
	"github.com/CyclopsV/cities-informer-skillbox/pkg/services"
	"log"
	"net/http"
)

func main() {
	go func() {
		log.Println("Запуск сервера")
		r := api.CreateRoutes()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			panic(err)
		}
	}()

	services.Closer(&api.Cities)
}
