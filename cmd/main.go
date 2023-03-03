package main

import (
	"champ"
	"log"
	"net/http"
)

func main() {
	c := champ.New()
	c.Route("GET", "/user", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello from user")
	})
	c.ListenAndServe(":8000")
}
