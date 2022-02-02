package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	fmt.Println("filenames!!!", filenames)
	var files []string
	for _, file := range filenames {
		fmt.Println("file", fmt.Sprintf("app/views/templates/%s.html", file))
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	fmt.Println("templates!!!", templates)
	templates.ExecuteTemplate(writer, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static", files))
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
