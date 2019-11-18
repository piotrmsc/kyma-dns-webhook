package internal

import (
	"net/http"
	"fmt"
	)

func RunServer(){
	mux := http.NewServeMux()

	mux.Handle("/present", PresentHandler())
	mux.Handle("/cleanup", CleanupHandler())

	svr := &http.Server{
		Addr: fmt.Sprintf(":%d", 9096),
		Handler:mux,
	}

	fmt.Println(svr.ListenAndServe())
}