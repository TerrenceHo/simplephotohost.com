package controllers

import (
	"fmt"
	"net/http"
)

func (g *Gallery) New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Gallery")
}

type Gallery struct {
	Photo string
}

func NewGallery() *Gallery {
	return &Gallery{
		Photo: "Hey",
	}
}
