package main

import (
	"bytes"
	"log"
	"net/http"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		f, _ := Htmls.Open(t.filename)
		buf := bytes.Buffer{}
		buf.ReadFrom(f)
		t.templ = template.Must(template.New(t.filename).Parse(buf.String()))
		log.Println("hellowold")
	})
	t.templ.Execute(w, nil)
}
