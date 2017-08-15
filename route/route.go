package route

import (
	"net/http"
	"github.com/go-xorm/xorm"
	"../model"
	_"io/ioutil"
	_"fmt"
)

func Function(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	m := req.Method
	if m == "GET" {
		model.GetAllFunction(res, req, engine)
	}else if m == "POST" {
		model.AddFunction(res, req, engine)
	}
}

func License(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	res.Write([]byte("License Page"))
	m := req.Method
	if m == "GET" {
		model.GetAllLicense(res, req, engine)
	}else if m == "POST" {
		model.AddLicense(res, req, engine)
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



