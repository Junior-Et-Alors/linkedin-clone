package utils

import (
	"linkedin-clone/src/entity"
	"net/http"
	"strconv"
)

func JsonEncoderMessage(tab []string, user entity.User, w http.ResponseWriter) {

	if len(user.GetEmail()) > 0 {
		tab = append(tab, user.GetEmail())
	}

	if len(user.GetFirstName()) > 0 {
		tab = append(tab, user.GetFirstName())
	}

	if len(user.GetLastname()) > 0 {
		tab = append(tab, user.GetLastname())
	}

	if len(user.GetPassword()) > 0 {
		tab = append(tab, string(user.GetPassword()))
	}

	if user.GetAge() != 0 {
		//convert age int to string for respect type tab array
		age := strconv.Itoa(user.GetAge())
		tab = append(tab, age)
	}

	if user.GetId() != 0 {
		//convert age int to string for respect type tab array
		id := strconv.Itoa(user.GetId())
		tab = append(tab, id)
	}

	JsonEncoderString(tab, w)
}

func SendReponseLoginIsOk(outputJsonResponse []int, w http.ResponseWriter) {
	outputJsonResponse = append(outputJsonResponse, 200)
	JsonEncoderInt(outputJsonResponse, w)
}

func SendReponseError(outputMessageError []int, msg int, w http.ResponseWriter) {
	outputMessageError = append(outputMessageError, msg)
	JsonEncoderInt(outputMessageError, w)
}
