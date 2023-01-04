package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var hostname string
var err error

func init() {
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
		io.WriteString(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Get Hostname</title>
		</head>
		<body>
			
			<h3>My Hostname is ` +  c.Value + ` </h3>
		
		</body>
		</html>`)
		return
	}

	c = &http.Cookie{
		Name:   "hsn",
		Value:  hostname,
		MaxAge: 7200,
	}

	http.SetCookie(w, c)

	io.WriteString(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Get Hostname</title>
		</head>
		<body>
			
			<h3>My Hostname is ` +  c.Value + ` </h3>
		
		</body>
		</html>`)

}