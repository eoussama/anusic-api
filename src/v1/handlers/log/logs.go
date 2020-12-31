package log

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/eoussama/anusic-api/src/shared/utils"
)

// Logs handles the logs request (/api/v1/logs/)
func Logs(w http.ResponseWriter, r *http.Request) {

	// Initializing the logIds list
	var logIDs []string

	// Constructing the logs directory path
	absPath, _ := filepath.Abs(".")
	path := filepath.Join(absPath, utils.LogDirectory)

	// Listing the logs directory
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		fileName := file.Name()
		logIDs = append(logIDs, fileName[:strings.Index(fileName, ".")])
	}

	// Returning the response
	utils.ReturnResponse(w, logIDs, nil)
}
