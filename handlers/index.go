package handlers

import (
	"net/http"

	"fmt"

	"github.com/GolangAce/experiment/common"
)

//IndexHandler handles index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started IndexHandler")
	path := r.URL.Path
	common.RenderTemplate(w, "templates/home.html", path)
	fmt.Println("Closed IndexHandler")
}
