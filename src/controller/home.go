package controller

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	fmt.Println("le serveur est actif")
	//fmt.Println(r.Host, r.URL.Path)
	//u := entity.User{Email: "coucou@gmail.com", Firstname: "homepage"}

	// entity := []entity.User{u}

	// var outputJsonResponse []string
	// for _, e := range entity {
	// 	utils.JsonEncoderString(outputJsonResponse, e, w)
	// }
}
