package utils

import (
	"encoding/json"
	"net/http"
)

func JsonEncoderInt(tab []int, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(tab)
}

func JsonEncoderString(tab []string, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(tab)
}
