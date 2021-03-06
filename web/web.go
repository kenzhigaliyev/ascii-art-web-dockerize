package student

import (
	"net/http"
	"strings"
	student "student/ascii"
	"text/template"
)

type ASCII struct {
	Str    string
	Font   string
	Output string
}

var Result ASCII

//GET method part
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.Method == "GET" {
		http.Error(w, "400 Not Found", http.StatusNotFound)
		return
	}

	if r.URL.Path == "/" && r.Method != "GET" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	val, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	val.ExecuteTemplate(w, "index", nil)
}

//POST method part
func Body(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii-art/" || (r.URL.Path != "/ascii-art/" && r.Method == "POST") {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" && r.URL.Path == "/ascii-art/" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	val, err := template.ParseFiles("templates/web.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.FormValue("str") == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	Result.Str = strings.ReplaceAll(r.FormValue("str"), "\n", "")
	if !student.Check(Result.Str) {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if r.FormValue("submit_btn") != "Submit" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	Result.Font = r.FormValue("font")

	if r.FormValue("font") == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	if Result.Font == "thinkertoy" || Result.Font == "standard" || Result.Font == "shadow" {
		Result.Output = student.Output(Result.Str, Result.Font)
		if Result.Output == "" {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		val.ExecuteTemplate(w, "web", Result)
	} else {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
}

//main function
func MainFunc() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/ascii-art/", Body)
	http.ListenAndServe(":7777", nil)
}
