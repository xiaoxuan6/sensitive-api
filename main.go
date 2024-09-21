package main

import (
	"github.com/xiaoxuan6/sensitiveCheck/router"
	"github.com/xiaoxuan6/sensitiveCheck/services"
	"net/http"
)

func main() {
	services.InitSensitive()
	go services.WaterDict()

	r := router.Register()
	_ = http.ListenAndServe(":9210", r)
}
