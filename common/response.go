package common

import (
	"encoding/json"
	"net/http"
)

func H(w http.ResponseWriter, data map[string]interface{}) {
	if _, ok := data["msg"]; !ok {
		code := data["code"].(int)
		data["msg"] = http.StatusText(code)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func HSuccess(w http.ResponseWriter) {
	data := make(map[string]interface{})
	data["code"] = http.StatusOK

	H(w, data)
}

func HSuccessWithData(w http.ResponseWriter, data interface{}) {
	item := make(map[string]interface{})
	item["code"] = http.StatusOK
	item["data"] = data

	H(w, item)
}

func HError(w http.ResponseWriter) {
	data := make(map[string]interface{})
	data["code"] = http.StatusBadRequest

	H(w, data)
}

func HErrorWithMsg(w http.ResponseWriter, msg string) {
	data := make(map[string]interface{})
	data["code"] = http.StatusBadRequest
	data["msg"] = msg

	H(w, data)
}
