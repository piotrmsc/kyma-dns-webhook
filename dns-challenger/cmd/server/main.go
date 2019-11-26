package main

import (
	"github.com/kyma-incubator/kyma-dns-webhook/dns-proxy/server"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	gcpProject := os.Getenv("GCE_PROJECT")

	if gcpProject == "" {
		log.Fatal("GCP Project not provided. Please provide env variables GCE_PROJECT")
	}

	gcpSA := os.Getenv("GCE_SERVICE_ACCOUNT_FILE")

	if gcpSA == "" {
		log.Fatal("GCP Service Account file path not provided. Please provide env variables GCE_SERVICE_ACCOUNT_FILE")
	}

	server.RunServer()
}