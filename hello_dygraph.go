package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var myString string = "Date, Temp\n"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("dygraph.html")
	if err != nil {
		log.Fatal("Could not parse template file: dygraph,html")
	}
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	myTime := time.Now()
	tmpString := myString
	for i := 0; i < 100; i++ {
		tmpString = tmpString + fmt.Sprintf("%s, %d\n", myTime.Add(time.Hour*time.Duration(24*i)).Format("01/02/2006"), i*2)
	}
	t.Execute(w, tmpString)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
