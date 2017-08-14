package route

import (
	"net/http"
	"github.com/go-xorm/xorm"
)


func License(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	res.Write([]byte("License Page"))
	m := req.Method
	if m == "GET" {

	}else if m == "POST" {

	}
}

func Config(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	res.Write([]byte("Config Page"))
	m := req.Method
	if m == "GET" {

	}else if m == "POST" {
		
	}
}

func Verify(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	res.Write([]byte("Verify Page"))
	m := req.Method
	if m == "GET" {

	}else if m == "POST" {
		
	}
}



