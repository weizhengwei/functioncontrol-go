package util

import (
    "log"
    "os"
    "sync"
    "io"
)

type mylogger struct {
    filename string
    *log.Logger
}

var logger *mylogger
var once sync.Once

// start loggeando
func GetLoggerInstance() *mylogger {
    once.Do(func() {
        logger = createLogger("app.log")
    })
    return logger
}

func createLogger(fname string) *mylogger {
    file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

    writers := []io.Writer{
            file,
            os.Stdout,
        }
    fileAndStdoutWriter := io.MultiWriter(writers...)
    return &mylogger{
        filename: fname,
        Logger:   log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile),
    }
}