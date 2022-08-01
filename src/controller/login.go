package controller

import (
	"bytes"
	"io"
	"io/ioutil"
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

	if e == "" || p == "" {
		msg := 400
		var outputMessageError []int
		utils.SendReponseError(outputMessageError, msg, w)
		return
	}
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

	u := entity.User{Email: e, Mdp: p}
	checkLogUser := usecase.Login(u, p)

	if !checkLogUser {
		msg := 403
		var outputMessageError []int
		utils.SendReponseError(outputMessageError, msg, w)
		return
	}

	var outputJsonResponse []int

	utils.SendReponseLoginIsOk(outputJsonResponse, w)

}
