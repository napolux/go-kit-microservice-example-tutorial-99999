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

	"napodate"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080",
			"http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our napodate service
	srv := napodate.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	getEndpoint := napodate.MakeGetEndpoint(srv)
	statusEndpoint := napodate.MakeStatusEndpoint(srv)
	validateEndpoint := napodate.MakeValidateEndpoint(srv)

	endpoints := napodate.Endpoints{
		GetEndpoint:      getEndpoint,
		StatusEndpoint:   statusEndpoint,
		ValidateEndpoint: validateEndpoint,
	}

	// HTTP transport
	go func() {
		log.Println("napodate is listening on port", *httpAddr)
		handler := napodate.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
