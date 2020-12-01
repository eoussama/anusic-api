package utils

import (
	"log"
	"os"
)

// file path of our logs
var path = "logs/logs.txt"

//setting up logs type
func Log(txt string, typ int) {
	if err := ensureDir("logs"); err != nil {
		os.Exit(1)
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	switch typ {
	case 0:
		log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile).Print(txt)
	case 1:
		log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile).Print(txt)
	default:
		log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile).Print(txt)
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
