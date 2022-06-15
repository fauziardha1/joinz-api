package main

import (
	"fmt"
	"joinz/pkg/api"
	"log"
	"net/http"
	"os"
)

func main() {
	router := api.SetupRouter()

	log.Print("We are up and running")
	port := os.Getenv("PORT")
	err := http.ListenAndServe(fmt.Sprintf("%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}

}
