package utilities

import (
	"log"
)

func Error(err error, msg ...string) {

	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
