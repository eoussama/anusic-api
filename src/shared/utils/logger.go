package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// LogDirectory is the name of the Logging dump location
var LogDirectory = "logs"

// InfoLogger hosts the logger that logs info
var InfoLogger *log.Logger

// WarningLogger hosts the logger that logs warnings
var WarningLogger *log.Logger

// ErrorLogger hosts the logger that logs errors
var ErrorLogger *log.Logger

// RequestLogger hosts the logger that logs requests
var RequestLogger *log.Logger

// init initializes the logger
func init() {

	// Constructing the log file name
	timestamp := time.Now().Unix()
	logFile := strconv.Itoa(int(timestamp)) + ".log"

	// Constructing the logs file path
	absPath, _ := filepath.Abs(".")
	directoryPath := filepath.Join(absPath, LogDirectory)
	filePath := filepath.Join(directoryPath, logFile)

	// Ensuring the logs directory
	os.Mkdir(directoryPath, 0777)

	// Opening the logging file
	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// Initializing the loggers
	InfoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "[WARNING] ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime)
	RequestLogger = log.New(file, "[REQUEST] ", log.Ldate|log.Ltime)

	// Setting both the console and the logging file as output targets
	InfoLogger.SetOutput(io.MultiWriter(os.Stdout, file))
	WarningLogger.SetOutput(io.MultiWriter(os.Stdout, file))
	ErrorLogger.SetOutput(io.MultiWriter(os.Stdout, file))
	RequestLogger.SetOutput(io.MultiWriter(os.Stdout, file))
}

// Log handles logging
func Log(input interface{}, logType int) {
	switch logType {
	case 1:
		WarningLogger.Println(input)
	case 2:
		ErrorLogger.Println(input)
	case 3:
		RequestLogger.Println(input)
	default:
		InfoLogger.Println(input)
	}
}
