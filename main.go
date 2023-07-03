package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ibilalkayy/backend/api"
)

func Execute() error {
	api.SetupRoutes()
	fmt.Println("Starting the server at: 8080")
	return http.ListenAndServe(":8080", nil)
}

func main() {
	if err := Execute(); err != nil {
		log.Fatal(err)
	}
}
