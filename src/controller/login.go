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
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
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

	e := contact["Email"]
	p := contact["Password"]

	if len(e) == 0 || len(p) == 0 {
		msg := 403
		var outputMessageError []int
		utils.SendReponseError(outputMessageError, msg, w)
		return
	}

	checkEmail := utils.CheckEmail(e)
	if !checkEmail {
		msg := 403
		var outputMessageError []int

		utils.SendReponseError(outputMessageError, msg, w)
		return
	}

	//password := []byte(p)
	//hashedpassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	u := entity.User{Email: e, Mdp: p}
	clientOptions := mongo.Initializer()
	checkLogUser := usecase.Login(clientOptions, u, p)

	if !checkLogUser {
		msg := 403
		var outputMessageError []int
		utils.SendReponseError(outputMessageError, msg, w)
		return
	}

	var outputJsonResponse []int

	utils.SendReponseLoginIsOk(outputJsonResponse, w)

}
