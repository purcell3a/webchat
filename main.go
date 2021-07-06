package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

type Topic struct{
	title string
	intro string
	route string
}

func GenerateTopic(t *Topic) string{
	switch t.title {
	case "topicOne":
		return t.route
	case "topicTwo":
		return t.route
	case "topicThree":
		return t.route
	}
	return "n"
}

/* handle a form, both the GET which displays the form
and the POST which processes it.*/
// func FormServer(w http.ResponseWriter, request *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// }

/* SIMPLE GET REQUEST */
func TopicOne(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "TopicOne", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}

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

