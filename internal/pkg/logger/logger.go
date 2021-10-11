package logger

import (
	"fmt"
	"log"
	"os"
)

// opt 使每一层使用自己的 logger, 实现 logger 层

var logger *log.Logger
var skip = 2

func Init() {
	initLogger()
}

func Debug(v ...interface{}) {
	debugWithSkip(skip, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	debugWithSkip(skip, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	infoWithSkip(skip, fmt.Sprint(v...))

}

func Infof(format string, v ...interface{}) {
	infoWithSkip(skip, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	warnWithSkip(skip, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	warnWithSkip(skip, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	errorWithSkip(skip, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	errorWithSkip(skip, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	fatalWithSkip(skip, fmt.Sprint(v...))

}

func Fatalf(format string, v ...interface{}) {
	fatalWithSkip(skip, fmt.Sprintf(format, v...))
}

func Panic(v ...interface{}) {
	panicWithSkip(skip, fmt.Sprint(v...))

}

func Panicf(format string, v ...interface{}) {
	panicWithSkip(skip, fmt.Sprintf(format, v...))
}

// todo mod
func debugWithSkip(skip int, msg string) {
	log.Println("[DEBUG]: ", msg)
}

func infoWithSkip(skip int, msg string) {
	log.Println("[INFO]: ", msg)
}

func warnWithSkip(skip int, msg string) {
	log.Println("[WARN]: ", msg)
}

func errorWithSkip(skip int, msg string) {
	log.Println("[ERROR]: ", msg)
}

func fatalWithSkip(skip int, msg string) {
	log.Println("[FATAL]: ", msg)
}

func panicWithSkip(skip int, msg string) {
	log.Println("[PANIC]: ", msg)
	panic(msg)
}

// todo mod
func initLogger() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}
