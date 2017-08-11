package route

import (
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Home Page"))
}

func Doc(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Doc Page"))
}