package route

import (
	"net/http"
	"fmt"
	"github.com/go-xorm/xorm"
)

func Home(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	res.Write([]byte("Home Page"))
	fmt.Println(engine)
}

