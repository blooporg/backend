package user

import (
	"fmt"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Sign In Page!! </h1>")
}
