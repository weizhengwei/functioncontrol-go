package main

import (
	"fmt"
	"net/http"
	"os"
	"./route"
	"./model"
	"./util"
	_ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
   
)

var engine *xorm.Engine
var BIND_ADDR = "localhost:9090"

func initxorm() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:r00t@/www?charset=utf8")
	if err != nil {
		fmt.Println("NewEngine failed", err)
		os.Exit(1)
	}

	f, err1 := os.Create("sql.log")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	engine.ShowSQL(true)
	engine.SetLogger(xorm.NewSimpleLogger(f))
	
	// 同步结构体与数据表
	if err2 := engine.Sync2(new(model.TbLicenseModule), new(model.TbFunctionModule)); err2 != nil {
		fmt.Printf("Fail to sync database: %v\n", err2)
	}
	fmt.Println("init OK")
}

func Home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Home Page"))
	util.GetLoggerInstance().Println("Home")
	fmt.Println("Home fmt")
}

func Doc(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Doc Page"))
	util.GetLoggerInstance().Println("Doc")
	fmt.Println("Doc fmt")
}

func HandleFunction(res http.ResponseWriter, req *http.Request) {
	route.Function(res, req, engine)
}

func HandleLicense(res http.ResponseWriter, req *http.Request) {
	route.License(res, req, engine)
}

func HandleConfig(res http.ResponseWriter, req *http.Request) {
	route.Config(res, req, engine)
}

func HandleVerify(res http.ResponseWriter, req *http.Request) {
	route.Verify(res, req, engine)
}

func main() {
	initxorm()
	http.HandleFunc("/", Home)
	http.HandleFunc("/doc", Doc)
	http.HandleFunc("/api/function", HandleFunction)
	http.HandleFunc("/api/license", HandleLicense)
	http.HandleFunc("/api/config", HandleConfig)
	http.HandleFunc("/api/verify", HandleVerify)

	util.GetLoggerInstance().Println("Server Start At", BIND_ADDR)
	err := http.ListenAndServe(BIND_ADDR, nil)
	if err != nil {
		util.GetLoggerInstance().Println("Start Server Failed:", err)
	}
}
