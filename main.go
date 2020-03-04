package main

import (
	"fmt"
	"net/http"

	// "github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/views"
)

var homeView, contactView, faqView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>***404 Page Not Found!***</h1>")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var nf http.Handler = http.HandlerFunc(pageNotFound)

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.NotFoundHandler = nf
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", r)
}
