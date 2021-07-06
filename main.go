package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	// "github.com/purcell3a/webchat/models/users"
	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func TopicOne(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "TopicOne", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}

// type templateHandler struct {
// 	once     sync.Once
// 	filename string
// 	templ    *template.Template
// }

// // this method loads the source file, compiles the template, executes it and
// //writes the output to the specifeid httm.ResponsWriter method bc sercehttp method
// //satisfies the http.handler interface and we can pass it directly to http.Handler
// func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	t.once.Do(func() {
// 		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
// 	})
// 	t.templ.Execute(w, nil)
// }

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/employeedb")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()

	// root
	// http.Handle("/", &templateHandler{filename: "chat.html"})
	http.HandleFunc("/topicone", TopicOne)
	http.HandleFunc("/", Index)
	// tmpl := template.Must(template.ParseFiles("chat.html"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// data := ""
	// 	tmpl.Execute(w, nil)
	// })
	// start the webserver
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	http.ListenAndServe(":3000", nil)
}

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
