package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Name struct {
	Name string
}

func hello(res http.ResponseWriter, req *http.Request) {

	content := genIndx(req)

	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, string(content))
}

func genIndx(req *http.Request) []byte {
	tmpl, err := template.ParseFiles("./cmd/syntax_task6/http/static/template.html")
	if err != nil {
		log.Fatal(err)
	}

	data := getName(req)
	f, err := os.OpenFile("./cmd/syntax_task6/http/static/index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile("./cmd/syntax_task6/http/static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func getName(req *http.Request) Name {
	name := req.URL.Query().Get("name")
	if name == "" { name = "Go" }
	return Name {
		Name: name,
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
