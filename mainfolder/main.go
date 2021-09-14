package main

import (
	"fmt"
	"html/template"
	"net/http"
	"student"
)

type ASCII struct {
	Str    string
	Font   string
	Output string
}

var Result ASCII

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	val, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	val.ExecuteTemplate(w, "index", nil)
}

func body(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" && r.URL.Path == "/ascii-art/" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ascii-art/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	val, err := template.ParseFiles("templates/web.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "500 Internal Server Error", http.StatusInternalServerError)
	}

	if r.FormValue("str") == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	Result.Str = r.FormValue("str")
	Result.Font = r.FormValue("font")
	Result.Output = student.Output(Result.Str, Result.Font)
	val.ExecuteTemplate(w, "web", Result)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii-art/", body)
	http.ListenAndServe(":7777", nil)

}

func main() {
	handleFunc()
}
