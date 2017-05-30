package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/GolangAce/experiment/common"
)

type Article struct {
	Title string
	Body  template.HTML
	Slug  string
}

//PostsHandler is handling the index page
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started postshandler")
	slug := r.URL.Path + "/"

	//Clone here because we want to see our writes
	session := common.Dial().Clone()
	defer session.Close()

	collection := session.DB("testdb").C("posts")

	var result Article

	err1 := collection.Find(bson.M{"slug": slug}).One(&result)
	if err1 != nil {
		http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
	}

	fmt.Println(result)
	common.RenderTemplate(w, "templates/posts.html", result)
	fmt.Println("Closed postshandler")
}
