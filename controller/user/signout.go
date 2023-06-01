package user

import (
	"fmt"
	"net/http"
)

func Signout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Sign Out Page!! </h1>")
}
