package handlers

import (
	"net/http"

	"github.com/GolangAce/experiment/common"
)

//CreditsHandler handles index page
func CreditsHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	common.RenderTemplate(w, "templates/credits.html", path)
}
