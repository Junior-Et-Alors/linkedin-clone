package main

import (
	"fmt"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenue sur la page login")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/ins", Register)

	http.ListenAndServe(":8080", nil)

}
