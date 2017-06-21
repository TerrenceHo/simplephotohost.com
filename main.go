package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>This site is driven by super awesome machine learning</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for.  Oops! Try again maybe?</h1>")
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name")) // ps httprouter.Params looks for URL parameter by :name and prints it out
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello) // Passes in Hello as a function, :name is a variable URL parameter

	// http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", router) // Listens and serves on localhost:3000, using router as a function
}
