package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:generate go-assets-builder -v Htmls -s="/resources/html" -o htmls.go ../resources/html
//go:generate go-assets-builder -v Images -s="/resources" -o images.go ../resources/image

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
	f, _ := Htmls.Open("/index.html")
	buf := bytes.Buffer{}
	buf.ReadFrom(f)
	tpl, _ := template.New("mytemp").Parse(buf.String())
	tpl.Execute(w, nil)
	//if err := t.ExecuteTemplate(w, "index.html", str); err != nil {
	//	log.Fatal(err)
	//}
}

func main() {
	fmt.Println("Open http://localhost:3000/")
	fmt.Println("Press ctrl+c to stop")

	mux := http.NewServeMux()
	//mux.Handle("/", http.FileServer(Htmls))
	mux.HandleFunc("/", htmlHandler0)
	//mux.Handle("/", &templateHandler{filename: "/index.html"})

	mux.Handle("/image/", http.FileServer(Images))

	if err := http.ListenAndServe(":3000", mux); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
