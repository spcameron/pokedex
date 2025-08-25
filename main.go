package main

import "net/http"

type Client struct {
	httpClient http.Client
}

func main() {
	config := &config{
		apiClient: Client{
			httpClient: http.Client{},
		},
	}

	startREPL(config)
}
