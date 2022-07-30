package controller

import (
	"bytes"
	"io"
	"io/ioutil"
	"linkedin-clone/db/mongo"
	"linkedin-clone/src/entity"
	"linkedin-clone/src/usecase"
	"linkedin-clone/src/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	r.Body = io.NopCloser(bytes.NewReader(body))

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	contact := make(map[string]string)

	contact["Email"] = r.PostFormValue("Email")
	contact["Password"] = r.PostFormValue("Password")
	contact["Firstname"] = r.PostFormValue("Firstname")
	contact["Lastname"] = r.PostFormValue("Lastname")

	e := contact["Email"]
	p := contact["Password"]
	f := contact["Firstname"]
	l := contact["Lastname"]

	if e == "" || p == "" || f == "" || l == "" {
		msg := 400
		var outputMessageError []int

		utils.SendReponseError(outputMessageError, msg, w)
		return
	}

	if len(e) == 0 || len(p) == 0 || len(f) == 0 || len(l) == 0 {
		msg := 403
		var outputMessageError []int

		utils.SendReponseError(outputMessageError, msg, w)
		return
	}

	email := utils.CheckEmail(e)
	if !email {

		msg := 403
		var outputMessageError []int

		utils.SendReponseError(outputMessageError, msg, w)
		return

	}

	password := []byte(p)
	hashedpassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	u := entity.User{Email: e, Password: hashedpassword, Firstname: f, Lastname: l}

	clientOptions := mongo.Initializer()
	check := usecase.CheckIsEmailExist(clientOptions, u)

	if check == true {
		msg := 403
		var outputMessageError []int

		utils.SendReponseError(outputMessageError, msg, w)
		return
	}
	go usecase.InsertUser(clientOptions, u)

	var outputJsonResponse []int

	utils.SendReponseLoginIsOk(outputJsonResponse, w)
}
