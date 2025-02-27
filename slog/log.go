package slog

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var log *Logger

//Logger 异步日志
type Logger struct {
	console bool
	warn    bool
	info    bool
	tformat func() string
	file    chan interface{}
	fileDB  chan interface{}
}

//NewLog 创建异步日志
func NewLog(level string, console bool, buf int) (*Logger, error) {
	log = &Logger{console: console, tformat: format}
	File, _ := os.Create("server.log")
	FileDB, _ := os.OpenFile("db.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if File != nil {
		FileInfo, err := File.Stat()
		if err != nil {
			return nil, err
		}
		mode := strings.Split(FileInfo.Mode().String(), "-")
		if strings.Contains(mode[1], "w") {
			strChan := make(chan interface{}, buf)
			log.file = strChan
			go func() {
				for {
					fmt.Fprintln(File, <-strChan)
				}
			}()
			defer func() {
				for len(strChan) > 0 {
					time.Sleep(1e9)
				}
			}()
		} else {
			return nil, errors.New("can't write.")
		}
	}

	if FileDB != nil {
		FileInfo, err := File.Stat()
		if err != nil {
			return nil, err
		}
		mode := strings.Split(FileInfo.Mode().String(), "-")
		if strings.Contains(mode[1], "w") {
			strChan := make(chan interface{}, buf)
			log.fileDB = strChan
			go func() {
				for {
					fmt.Fprintln(FileDB, <-strChan)
				}
			}()
			defer func() {
				for len(strChan) > 0 {
					time.Sleep(1e9)
				}
			}()
		} else {
			return nil, errors.New("can't write.")
		}
	}
	switch level {
	case "Warn":
		log.warn = true
		return log, nil
	case "Info":
		log.warn = true
		log.info = true
		return log, nil
	}
	return nil, errors.New("level must be Warn or Info.")
}

func Fatal(info ...interface{}) {
	if log.console {
		fmt.Println("Fatal", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Fatal", log.tformat(), info)
	}
	os.Exit(1)
}

//Error 错误级别
func Error(info ...interface{}) {
	if log.console {
		fmt.Println("Error", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Error", log.tformat(), info)
	}
}
func ErrorDB(info ...interface{}) {
	if log.console {
		fmt.Println("Error", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Error", log.tformat(), info)
	}
	if log.fileDB != nil {
		log.fileDB <- fmt.Sprintln("Error", log.tformat(), info)
	}
}

func Warn(info ...interface{}) {
	if log.warn && log.console {
		fmt.Println("Warn", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Warn", log.tformat(), info)
	}
}
func Info(info ...interface{}) {
	if log.info && log.console {
		fmt.Println("Info", log.tformat(), info)
	}
	if log.file != nil {
		log.file <- fmt.Sprintln("Info", log.tformat(), info)
	}
}
func Close() {
	for len(log.file) > 0 {
		time.Sleep(1e8)
	}
}
func format() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	return time.Now().Format("2006-01-02 15:04:05") + fmt.Sprintf(" %s:%d ", file, line)
}
