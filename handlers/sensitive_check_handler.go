package handlers

import (
	"github.com/xiaoxuan6/sensitiveCheck/common"
	request2 "github.com/xiaoxuan6/sensitiveCheck/request"
	"github.com/xiaoxuan6/sensitiveCheck/services"
	"net/http"
)

// Filter 直接移除敏感词
func Filter(w http.ResponseWriter, r *http.Request) {
	request, err := request2.Validate(r)
	if err != nil {
		common.HErrorWithMsg(w, err.Error())
		return
	}

	data := services.Filter.Filter(request.Content)
	common.HSuccessWithData(w, data)
}

// FindAll 筛选成所有的敏感词
func FindAll(w http.ResponseWriter, r *http.Request) {
	request, err := request2.Validate(r)
	if err != nil {
		common.HErrorWithMsg(w, err.Error())
		return
	}

	data := services.Filter.FindAll(request.Content)
	common.HSuccessWithData(w, data)
}

// Replace 把词语中的字符替换成指定的字符，这里的字符指的是rune字符
func Replace(w http.ResponseWriter, r *http.Request) {
	request, err := request2.Validate(r)
	if err != nil {
		common.HErrorWithMsg(w, err.Error())
		return
	}

	data := services.Filter.Replace(request.Content, '*')
	common.HSuccessWithData(w, data)
}

// Validate 验证内容是否ok，如果含有敏感词，则返回false和第一个敏感词。
func Validate(w http.ResponseWriter, r *http.Request) {
	request, err := request2.Validate(r)
	if err != nil {
		common.HErrorWithMsg(w, err.Error())
		return
	}

	ok, data := services.Filter.Validate(request.Content)
	common.HSuccessWithData(w, map[string]interface{}{"invalid": ok, "sensitive_word": data})
}
