package handlers

import (
	"github.com/xiaoxuan6/sensitiveCheck/common"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	common.H(w, map[string]interface{}{"code": http.StatusNotFound})
}

func MethodNotAllow(w http.ResponseWriter, r *http.Request) {
	common.H(w, map[string]interface{}{"code": http.StatusMethodNotAllowed})
}
