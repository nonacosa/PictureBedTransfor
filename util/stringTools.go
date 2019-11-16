package util

import (
	"strings"
)

func UrlIsIamge(url string ) bool{
	if strings.Contains(url,"png") {
		return true
	}

	if strings.Contains(url,"jpg") {
		return true
	}

	if strings.Contains(url,"gif") {
		return true
	}

	return false
}
