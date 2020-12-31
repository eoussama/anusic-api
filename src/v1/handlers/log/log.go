package log

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"

	"github.com/eoussama/anusic-api/src/shared/models"
	"github.com/eoussama/anusic-api/src/shared/utils"
	"github.com/gorilla/mux"
)

// Log handles the logs request (/api/v1/logs/:id)
func Log(w http.ResponseWriter, r *http.Request) {

	// Getting the passed ID
	vars := mux.Vars(r)
	id := vars["id"]

	// Constructing the logs file path
	absPath, _ := filepath.Abs(".")
	path := filepath.Join(absPath, utils.LogDirectory, id+".log")

	logFile, err := os.Open(path)
	defer logFile.Close()

	// Returning the response
	if err != nil {
		utils.ReturnResponse(w, nil, models.Error.LogNotFound(models.Error{}, id))
	} else {
		logData := []string{}

		scanner := bufio.NewScanner(logFile)
		for scanner.Scan() {
			logData = append(logData, scanner.Text())
		}

		utils.ReturnResponse(w, logData, nil)
	}
}
