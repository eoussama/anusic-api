package utils

import (
	"log"
	"os"
)

var WarningLogger *log.Logger
var InfoLogger *log.Logger
var ErrorLogger *log.Logger

// file path of our logs
var path = "logs/logs.log"

func init() {
	if err := ensureDir("logs"); err != nil {
		os.Exit(1)
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

//setting up logs type
func Log(txt string, typ int) {
	switch typ {
	case 0:
		InfoLogger.Print(txt)
	case 1:
		WarningLogger.Print(txt)
	default:
		ErrorLogger.Print(txt)
	}

}

//Create directory if doesn't exist
func ensureDir(dirName string) error {

	err := os.Mkdir(dirName, os.ModeDir)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
