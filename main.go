package main

import (
	"fmt"
	"github.com/kyma-incubator/kyma-dns-webhook/internal"
)

func main() {
	fmt.Println("hello there!")
	internal.RunServer()
}