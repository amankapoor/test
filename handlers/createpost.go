package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GolangAce/experiment/common"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName         = "testdb"
	collectionName = "posts"
)

//CreatePostHandler maintains admin zone

type Posts struct {
	//meta data about posts
	ID bson.ObjectId `bson:"_id,omitempty"`
	/*PublishedTimeStamp time.Time
	UpdatedTimeStamp   time.Time
	Status             string*/
	Slug string `bson:"slug"`
	//Views              string
	//data about posts
	Title string `bson:"Title"`
	/*Subheading    string
	Author        string*/
	Body template.HTML `bson:"Body"`
	//Topics string
	//FeaturedImage string
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("started createpost handler")
	common.RenderTemplate(w, "templates/createPost.html", nil)

	if r.Method == "POST" {
		title := r.FormValue("title")
		slug := r.FormValue("slug")
		body := template.HTML(r.FormValue("body"))

		id := bson.NewObjectId()
		//Inserting
		p := Posts{
			ID:    id,
			Title: title,
			Slug:  slug,
			Body:  body,
		}

		//We don't care about reading our writes here, so copy
		session := common.Dial().Copy()
		defer session.Close()
		collection := session.DB("testdb").C("posts")

		insertionError := collection.Insert(p)
		if insertionError != nil {
			panic(insertionError)
		}

		fmt.Println(id.Hex())
	}
	fmt.Println("closed createpost handler")
}
