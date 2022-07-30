package main

import (
	"linkedin-clone/config/info"
	infoLogger "linkedin-clone/logger"
	"linkedin-clone/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := ":8080"

	r := mux.NewRouter()

	r.HandleFunc("/", controller.HomePage)
	r.HandleFunc("/login", controller.LoginPage).Methods("POST")

	r.HandleFunc("/register", controller.RegisterUser).Methods("POST")

	// profil := r.PathPrefix("/profil").Subrouter()
	// profil.HandleFunc("/parametre", controller.SettingPage)
	// profil.HandleFunc("/info/{name:[a-z]+}", controller.ProfilPage)

	go info.Server(r, port)
	infoLogger.Error(r)

	http.Handle("/", r)
	http.ListenAndServe(port, nil)

}
