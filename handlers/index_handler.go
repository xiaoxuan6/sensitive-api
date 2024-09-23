package handlers

import (
	"github.com/xiaoxuan6/sensitive-api/common"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	common.HSuccess(w)
}
