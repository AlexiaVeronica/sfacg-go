package request

import (
	"net/url"
)

type ConfigRequest struct {
	App      bool
	Cookie   string
	AppKey   string
	DeviceId string
	Host     string
}

var Request = ConfigRequest{
	App:      true,
	AppKey:   "FMLxgOdsfxmN!Dt4",
	DeviceId: "240a90cc-4c40-32c7-b44e-d4cf9e670605",
}

func Get(path string) *HttpUtils {
	return Req(path, "GET")
}

func Post(path string) *HttpUtils {
	return Req(path, "POST")
}
func Put(path string) *HttpUtils {
	return Req(path, "PUT")
}

func Req(path, method string) *HttpUtils {
	if Request.App {
		Request.Host = "https://api.sfacg.com/"
	} else {
		Request.Host = "https://minipapi.sfacg.com/pas/mpapi/"
	}
	return &HttpUtils{
		method:     method,
		query_data: &url.Values{},
		url:        Request.Host + path,
		Cookie:     Request.Cookie,
		Config:     Request,
	}
}
