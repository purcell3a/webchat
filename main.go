package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// this method loads the source file, compiles the template, executes it and
//writes the output to the specifeid httm.ResponsWriter method bc sercehttp method
//satisfies the http.handler interface and we can pass it directly to http.Handler
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// root
	http.Handle("/", &templateHandler{filename: "chat.html"})

	// start the webserver
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// const form = `<html><body><form action="#" method="post" name="bar">
// <input type="text" name="in"/>
// <input type="submit" value="Submit"/>
// </form></html></body>`

/* handle a simple get request */
// func SimpleServer(w http.ResponseWriter, request *http.Request) {
// 	io.WriteString(w, "<h1>hello, world</h1>")
// }

// func HelloServer(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("Inside HelloServer handler")
// 	fmt.Fprint(w, "Hello, "+req.URL.Path[1:])
// }

/* handle a form, both the GET which displays the form
and the POST which processes it.*/
// func FormServer(w http.ResponseWriter, request *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")

// 	switch request.Method {
// 	case "GET":
// 		/* display the form to the user */
// 		io.WriteString(w, form)
// 	case "POST":
// 		/* handle the form data, note that ParseForm must
// 		be called before we can extract form data with Form */
// 		// request.ParseForm();
// 		//io.WriteString(w, request.Form["in"][0])
// 		// easier method:
// 		io.WriteString(w, request.FormValue("in"))
// 	}
// }

// func main() {

// 	http.HandleFunc("/test1", SimpleServer)
// 	http.HandleFunc("/test2", FormServer)
// 	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
// 		panic(err)
// 	}
// }