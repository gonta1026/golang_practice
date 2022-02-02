package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "jello!", "layout", "top")
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// utils.LogFatalln(err)
	// t.Execute(w, "hello")
}
