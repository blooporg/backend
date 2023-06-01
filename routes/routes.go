package routes

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller"
	"github.com/ibilalkayy/backend/controller/loops"
	"github.com/ibilalkayy/backend/controller/user"
)

func Routes() {
	http.HandleFunc("/", controller.Index)

	// User pages
	http.HandleFunc("/signin", user.Signin)
	http.HandleFunc("/signup", user.Signup)
	http.HandleFunc("/signout", user.Signout)

	// Loops page
	http.HandleFunc("/loops", loops.Loops)
}
