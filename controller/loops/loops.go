package loops

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller/templates"
)

func Loops(w http.ResponseWriter, r *http.Request) error {
	return templates.LoopsTmpl.Execute(w, nil)
}
