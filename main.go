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
    "github.com/elgs/gostrgen"
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
}

func Doc(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Doc Page"))
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
	http.HandleFunc("/api/license", HandleLicense)
	http.HandleFunc("/api/config", HandleConfig)
	http.HandleFunc("/api/verify", HandleVerify)

	util.GetLoggerInstance().Println("Server Start At", BIND_ADDR)
	err := http.ListenAndServe(BIND_ADDR, nil)
	if err != nil {
		util.GetLoggerInstance().Println("Start Server Failed:", err)
	}
}

func GenerateRandomString() (string,error) {
	charsToGenerate := 16
	charSet := gostrgen.Upper | gostrgen.Digit
	includes := "" //"[]{}<>" // optionally include some additional letters
	excludes := "Ol"     //exclude big 'O' and small 'l' to avoid confusion with zero and one.

	str, err := gostrgen.RandGen(charsToGenerate, charSet, includes, excludes)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var ret string
	var k int
	for j := 0; j < 4; j++ {
		k = j*4
		if j < 3 {
			ret += str[k:k+4]+"-"
		}else{
			ret += str[k:k+4]
		}
	}
	return ret, nil
}