package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"os"
	"./route"
	"./model"
	_ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	logfile, err := os.OpenFile("test.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
    if err != nil {
        fmt.Printf("%s\r\n", err.Error())
        os.Exit(-1)
    }
    defer logfile.Close()
    writers := []io.Writer{
        logfile,
        os.Stdout,
    }
    fileAndStdoutWriter := io.MultiWriter(writers...)
    logger := log.New(fileAndStdoutWriter, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
    logger.Println("hello")
    logger.Println("oh....")
}

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

func home(res http.ResponseWriter, req *http.Request) {
	route.Home(res, req, engine)
}

func Doc(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Doc Page"))
}

func main() {
	initxorm()
	http.HandleFunc("/", home)
	http.HandleFunc("/doc", Doc)
	http.ListenAndServe(":9090", nil)
}