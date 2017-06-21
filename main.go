package main

import (
	"fmt"
	"net/http"

	// "github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Tpye", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name")) // ps httprouter.Params looks for URL parameter by :name and prints it out
//}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r) // Listens and serves on localhost:3000, using router as a function

	// router := httprouter.New()
	// router.GET("/hello/:name", Hello) // Passes in Hello as a function, :name is a variable URL parameter
}
