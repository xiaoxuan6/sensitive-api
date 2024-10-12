package main

import (
	"github.com/xiaoxuan6/sensitive-api/router"
	"github.com/xiaoxuan6/sensitive-api/services"
	"net/http"
)

func main() {
	services.InitSensitive()
	go services.WatcherDict()

	r := router.Register()
	_ = http.ListenAndServe(":9210", r)
}
