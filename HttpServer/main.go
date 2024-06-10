package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const serverPort = 3333

var serverLogger = log.New(os.Stdout, "server: ", log.Ldate|log.Ltime)
var clientLogger = log.New(os.Stdout, "client: ", log.Ldate|log.Ltime)

func setupServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serverLogger.Printf("%s /\n", r.Method)
		if r.Method == http.MethodPost {
			serverLogger.Printf("query id: %s\n", r.URL.Query().Get("id"))
			serverLogger.Printf("content-type: %s\n", r.Header.Get("content-type"))
			serverLogger.Printf("headers:\n")
			for headerName, headerValue := range r.Header {
				fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
			}
			reqBody, err := io.ReadAll(r.Body)
			if err != nil {
				serverLogger.Printf("could not read request body: %s\n", err)
			}
			serverLogger.Printf("request body: %s\n", reqBody)
		}
		fmt.Fprintf(w, `{"message": "hello!"}`)
		// time.Sleep(35 * time.Second)
	})
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			serverLogger.Printf("%s\n", err)
		}
	}
}

func setupGetRequest() {
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		clientLogger.Printf("%s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		clientLogger.Printf("%s\n", err)
		os.Exit(1)
	}

	clientLogger.Printf("got response!\n")
	clientLogger.Printf("status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		clientLogger.Printf("%s\n", err)
		os.Exit(1)
	}
	clientLogger.Printf("response body: %s\n", resBody)
}

func setupPostRequest() {
	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%d?id=1234", serverPort)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		clientLogger.Printf("%s\n", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		clientLogger.Printf("%s\n", err)
		os.Exit(1)
	}

	clientLogger.Printf("got response!\n")
	clientLogger.Printf("status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		clientLogger.Printf("%s\n", err)
		os.Exit(1)
	}
	clientLogger.Printf("response body: %s\n", resBody)
}

func main() {
	go setupServer()
	time.Sleep(100 * time.Millisecond)

	setupGetRequest()
	setupPostRequest()
}
