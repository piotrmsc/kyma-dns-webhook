package main

import (
	"github.com/kyma-incubator/kyma-dns-webhook/internal"
)

func main() {
	//TODO env handling
	//GCE_PROJECT & GCE_SERVICE_ACCOUNT_FILE for gcloud

	internal.RunServer()
}