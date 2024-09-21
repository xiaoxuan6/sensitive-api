package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Request struct {
	Content string `json:"content"`
}

func Validate(r *http.Request) (Request, error) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, err
	}

	if len(request.Content) < 1 {
		return request, errors.New("params content not empty")
	}

	return request, nil
}
