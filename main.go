package main

import (
	"fmt"
	"github.com/kyma-incubator/kyma-dns-webhook/internal"
	"os"
)

func main() {
	gcpProject := os.Getenv("GCE_PROJECT")

	if gcpProject == "" {
		//TODO error handling
		fmt.Println("GCP Project not provided. Please provide env variables GCE_PROJECT")
	}

	gcpSA := os.Getenv("GCE_SERVICE_ACCOUNT_FILE")

	if gcpSA == "" {
		//TODO error handling
		fmt.Println("GCP Service Account file path not provided. Please provide env variables GCE_SERVICE_ACCOUNT_FILE")
	}

	internal.RunServer()
}