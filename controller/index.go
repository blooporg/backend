package controller

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "<h1> Hello World!! </h1>")
	return nil
}
