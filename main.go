package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"os"
	"./route"
)

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

func main() {
	http.HandleFunc("/", route.Home)
	http.HandleFunc("/doc", route.Doc)
	http.ListenAndServe(":9090", nil)
}