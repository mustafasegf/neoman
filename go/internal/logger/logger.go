package logger

import (
	"fmt"
	"log"
	"os"
)

func createDirectory() {
	path, _ := os.Getwd()

	if _, err := os.Stat(fmt.Sprintf("%s/logs", path)); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}
}

func getLogWritter() *os.File {
	path, err := os.Getwd()
	if err != nil {
		log.Panic("get os wd", err.Error())
	}

	createDirectory()
	file, err := os.OpenFile(path+"/logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic("open file path", err.Error())
	}

	return file
}

func SetLogger() {
	log.SetOutput(getLogWritter())
	log.SetFlags(0)
}
