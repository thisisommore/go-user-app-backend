package util

import (
	"log"
	"testing"
)

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func HandleTestError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
