package main

import (
	"log"

	osa "github.com/samiam2013/openaistatusapi"
)

func main() {
	componentsStatus, err := osa.GetComponent(osa.API)
	if err != nil {
		log.Fatalf("Failed getting component status: %v", err)
	}
	log.Printf("%s %s", osa.API, componentsStatus)
}
