package main

import (
	"net/http"
	"os"
	"github.com/flynshuePersonal/testWebApp/views"
	"github.com/gorilla/mux"
)

var homeView *views.View

var myenv *envvars

type envvars struct {
	User     string
	Hostname string
}

func newEnvVars() *envvars {
	hostname := os.Getenv("HOSTNAME")
	username := os.Getenv("USER")
	return &envvars{
		User: username,
		Hostname: hostname,
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, myenv)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	myenv = newEnvVars()
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}
