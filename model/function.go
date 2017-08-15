package model

import (
	"time"
	"github.com/go-xorm/xorm"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type TbFunctionModule struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	FunctionNumber string    `xorm:"VARCHAR(32)"`
	FunctionType   int       `xorm:"INT(11)"`
	FunctionName   string    `xorm:"VARCHAR(255)"`
	State          int       `xorm:"default 0 INT(11)"`
	CreateTime     time.Time `xorm:"created"`
}

//curl localhost:9090/api/function
func GetAllFunction(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	results, err := engine.QueryString("select * from tb_function_module")
	if err != nil {
		fmt.Println(err)
		return
	}
	ret, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.Write(ret)
}

//curl -X POST -d "@addfunction.json" localhost:9090/api/function
func AddFunction(res http.ResponseWriter, req *http.Request, engine *xorm.Engine) {
	ret, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var functionmodule TbFunctionModule
	if err := json.Unmarshal(ret, &functionmodule); err != nil{
		fmt.Println(err)
		return
	}

	affected, err := engine.Insert(&functionmodule)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = affected
	res.Write([]byte("add functon ok"))
}