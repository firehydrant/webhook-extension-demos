package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	kc, err := NewKubeClient()
	if err != nil {
		log.Fatalf("could not generate kubernetes client: %w", err)
	}

	fmt.Println("starting server...")
	handler := NewWebhookHandler(kc)
	http.Handle("/webhooks", handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
