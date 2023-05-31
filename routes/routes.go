package routes

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller"
)

func Routes() {
	http.HandleFunc("/", controller.Index)
}
