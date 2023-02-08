package main

import (
	"github.com/CyclopsV/cities-informer-skillbox/api"
	"log"
	"net/http"
)

func main() {
	log.Println("Старт программы")

	r := api.CreateRoutes()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
