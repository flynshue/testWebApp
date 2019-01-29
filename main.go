package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

type data struct {
	Name     string
	Hostname string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	hostname := os.Getenv("HOSTNAME")
	myData := data{Name: "Felicia", Hostname: hostname}
	if err := homeTemplate.Execute(w, myData); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml",
		"views/layouts/footer.gohtml")
	if err != nil {
		log.Fatal("Template ParseFiles Error: ", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}
