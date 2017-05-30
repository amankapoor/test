package main

import (
	"net/http"

	"github.com/GolangAce/experiment/common"
	"github.com/GolangAce/experiment/handlers"
	"github.com/gorilla/mux"
)

/*//Article represents all the properties of an article
type Article struct {
	ArticleID          string
	Title              string
	Subheading         string
	Author             string
	PublishedTimeStamp time.Time
	UpdatedTimeStamp   time.Time
	Permalink          string
	FeaturedImage      string
	Body               string
	Views              string
}

//Member represets all the properties that a member has
type Member struct {
	MemberID     string
	MemberType   string
	FirstName    string
	LastName     string
	Twitter      string
	Facebook     string
	GooglePlus   string
	LinkedIn     string
	ProfilePhoto string
	ShortBio     string
	FullBio      string
}

//Image represents all the properties than an image has
type Image struct {
	ImageID            string
	URL                string
	Height             int
	Width              int
	IsInArticles       []string
	IsFeaturedImage    bool
	IsFeaturedImageFor []string
	Title              string
	AltTag             string
}*/

func main() {
	mainSession := common.Dial()
	defer mainSession.Close()
	common.Database("testdb")

	fs := http.FileServer(http.Dir("./static"))
	r := mux.NewRouter()
	//r.StrictSlash(true)
	http.Handle("/", r)
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/credits", handlers.CreditsHandler)
	//r.HandleFunc("/credits/", handlers.CreditsHandler)
	r.HandleFunc("/create", handlers.CreatePostHandler)
	//r.HandleFunc("/create/", handlers.CreatePostHandler)
	http.Handle("/static", http.StripPrefix("/static", fs))
	r.HandleFunc("/{[a-zA-Z0-9-]+}", handlers.PostsHandler)
	http.ListenAndServe(":8080", nil)
}

/* Issues with uncommented r.StrictSlash

1. When you go to localhost:8080, you will see that postshandler is also getting a call. This is extra calling, and only indexHandler should have it. But both are getting it, first postsHandler and then IndexHandler.

2. There are un-necessary extra callings, check console.

3. After functions return, you won't find a connection closed in terminal.
*/

/* Issues with commented r.StrictSlash

1. After functions return, you won't find a connection closed in terminal.

*/

/* My other issue is

1. I want to keep url scheme without trailing slash. But considering SEO, same link with and w/o trailing slash are different!  - https://webmasters.googleblog.com/2010/04/to-slash-or-not-to-slash.html

So, see the commented handlefuncs which have slash in end, should I be doing this, and put canonical if ends with trailing slash? Or we have some better approach?

2. In essence, I want only the url which has no traling slash to exist, but a user in RARE case may hit a link with trailing slash. I can't serve a 200 on both. I want to redirect it to the non-trailing slash version. So, what will the header of trailing slash version will serve - redirect 301 and the trailing one will have a rel canonical with trailing slash link?

If there is some other better approach, let me know.

3. I just want one of these to exist, that's why used r.StrictSlash, but it was messing things up.

4. And I want to see sessions being closed when handler returns in the mongod terminal? Or, maybe golang runs these handlers on new goroutines, that's why it is not closing? I am guessing it, because in this case, all goroutines will halt when main will return as I am not doing any sync.WaitGroup. Is this the case, because I feel so, as when I return main by pressing Ctrl+C in terminal, all the connections close.

I want to see proper closing.

Please help me solve all this Kamesh.

*/
