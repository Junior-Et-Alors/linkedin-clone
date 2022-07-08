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
	fmt.Fprintf(w, "hello word")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenue sur la page login")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "page coucou")
}

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".html"
// 	body, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/ins", Register)
	//http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)

}
