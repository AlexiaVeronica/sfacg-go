package request

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
)

var Request = Template.ConfigRequest{}

func Host() string {
	var host string
	if Request.App {
		host = "https://api.sfacg.com/"
	} else {
		host = "https://minipapi.sfacg.com/pas/mpapi/"
	}
	return host
}

func Get(path string) *HttpUtils {
	return NewHttpUtils(Host(), path, "GET", Request).Request()
}
func Post(path string) *HttpUtils {
	return NewHttpUtils(Host(), path, "POST", Request).Request()
}
func Put(path string) *HttpUtils {
	return NewHttpUtils(Host(), path, "PUT", Request).Request()
}
