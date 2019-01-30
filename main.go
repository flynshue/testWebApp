package main

import (
	"net/http"
	"os"
	"github.com/flynshuePersonal/testWebApp/views"
	"github.com/gorilla/mux"
)

var homeView *views.View

type data struct {
	User     string
	Hostname string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	hostname := os.Getenv("HOSTNAME")
	username := os.Getenv("USER")
	myData := data{User: username, Hostname: hostname}
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, myData)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}
