package internal

import (
	"fmt"
	"net/http"
	"github.com/go-acme/lego/v3/providers/dns/gcloud"

)

func PresentHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO validation

		gcpProvider, err := gcloud.NewDNSProvider()
		if err != nil {
			//TODO error handling
			fmt.Printf("could not get gcloud provider: %v", err)
		}

		err = gcpProvider.Present("mst.kyma-goat.ga", "tokien", "keyAuth")
		if err != nil {
			//TODO error handling
			fmt.Printf("present req failed: %v", err)
		}

		return
	})
}

func CleanupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gcpProvider, err := gcloud.NewDNSProvider()
		if err != nil {
			//TODO error handling
			fmt.Printf("could not get gcloud provider: %v", err)
		}

		err = gcpProvider.CleanUp("mst.kyma-goat.ga", "tokien", "keyAuth")
		if err != nil {
			//TODO error handling
			fmt.Printf("present req failed: %v", err)
		}

		return
	})
}
