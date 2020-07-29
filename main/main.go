package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	basicgokitfirst "basic-go-kit-first"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8000", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// service
	serv := basicgokitfirst.DateService{}
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoint

	endpoint := basicgokitfirst.Endpoint{
		GetEndpoint:      basicgokitfirst.MakeGetEndpoint(serv),
		StatusEndpoint:   basicgokitfirst.MakeStatusEndpoint(serv),
		ValidateEndpoint: basicgokitfirst.MakeValidateEndpoint(serv),
	}

	// Http transport

	go func() {
		log.Println("basic go kit is listening on port:", *httpAddr)
		handler := basicgokitfirst.NewHttpServer(ctx, endpoint)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)

}
