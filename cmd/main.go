package main

import (
	infoLogger "linkedin-clone/logger"
	"linkedin-clone/src/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/", controller.HomePage)
	r.HandleFunc("/login", controller.LoginPage).Methods("POST")
	r.HandleFunc("/register", controller.RegisterUser).Methods("POST")

	// profil := r.PathPrefix("/profil").Subrouter()
	// profil.HandleFunc("/parametre", controller.SettingPage)
	// profil.HandleFunc("/info/{name:[a-z]+}", controller.ProfilPage)

	//go info.Server(r, port)
	infoLogger.Error(r)

	http.Handle("/", r)
	http.ListenAndServe(port, nil)

}
