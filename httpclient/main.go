package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const server_port = 3000

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
			fmt.Fprintf(w, `{"message": "Hello!"}`)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", server_port),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)

	request_url := fmt.Sprintf("http://localhost:%d", server_port)
	req, err := http.NewRequest(http.MethodGet, request_url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	// res, err := http.Get(request_url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", res_body)
}
