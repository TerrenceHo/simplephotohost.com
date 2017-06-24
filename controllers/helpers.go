package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// parseForm is a general usecase helper function that takes in a request
// and a pointer to something.  It will attempt to ParseForm the request
// and if not possible return err.  Otherwise it will decode the Request
// and pass the information into the interface
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil { //error check on request
		return (err)
	}

	dec := schema.NewDecoder() //new decoder
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	} //decodes data into the SignupForm sturct

	return nil

}
