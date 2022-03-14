package utilities

import (
	"log"
)

// Error ErrorHandler is a function that handles errors
func Error(err error, msg ...string) {

	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
