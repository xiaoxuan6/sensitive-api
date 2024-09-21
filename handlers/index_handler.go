package handlers

import (
	"github.com/xiaoxuan6/sensitiveCheck/common"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	common.HSuccess(w)
}
