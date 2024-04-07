package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("Olá Mundo!"))

	u1 := usuario{
		Nome:  "Luís Felipe",
		Email: "lf@gmail.com",
	}

	templates.ExecuteTemplate(w, "home.html", u1)

}

func main() {

	templates = template.Must(template.ParseGlob("*html"))

	http.HandleFunc("/home", home)

	fmt.Println("Escutando na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}

// localhost:5000/home
