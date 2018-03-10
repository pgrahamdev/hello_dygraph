package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var csvString string = "Date, Temp\n"

const templateFile string = "dygraph.html"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Could not parse template file: " + templateFile)
	}

	myTime := time.Now()
	tmpString := csvString
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
