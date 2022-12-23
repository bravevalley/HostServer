package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var tpl *template.Template
var hostname string
var err error

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
	hostname, err = os.Hostname(); if err != nil {
		log.Fatalln("Unable to retrieve hostname")
	}
}

func main() {
	s := &http.Server{
		Addr:           ":18080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	http.HandleFunc("/", indexHandler)
	
	log.Fatalln(s.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("hsn")
	if err == nil {
		tpl.Execute(w, c.Value)
		return
	}

	c = &http.Cookie{
		Name:   "hsn",
		Value:  hostname,
		MaxAge: 7200,
	}

	http.SetCookie(w, c)

	tpl.Execute(w, c.Value)
}
