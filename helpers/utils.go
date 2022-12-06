package helpers

import (
	"log"
)

// Log erros
func LogError(err error) {
	log.Printf("Error - %v \n", err.Error())
}
