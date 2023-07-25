package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/firehydrant/demo-kube-restarter/extensions"
)

// WebhookHandler holds the handler logic for the webhook endpoint
type WebhookHandler struct {
	kubeClient *KubeClient
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(kubeClient *KubeClient) *WebhookHandler {
	return &WebhookHandler{kubeClient: kubeClient}
}

// sendMessage sends a HTTP POST request to the specified URL with the specified response body
func (h *WebhookHandler) sendMessage(url string, response *extensions.Response) error {
	jsonBody, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling request body: %w", err)
	}

	// Create a new HTTP request
	replyReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	// Set the request header Content-Type to application/json
	replyReq.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the HTTP request and get the HTTP response
	resp, err := client.Do(replyReq)
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status
	fmt.Printf("Response status: %v\n", resp.Status)

	return nil
}

// ServeHTTP is the HTTP handler for the "/webhooks" route
func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parse the JSON data in the request body
	var req extensions.Payload
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Failed to parse JSON request", http.StatusBadRequest)
		return
	}

	namespace, deploymentName := req.Data.CommandArguments[0], req.Data.CommandArguments[1]
	oldDeployment, newDeployment, err := h.kubeClient.restartDeployment(namespace, deploymentName)
	if err != nil {
		http.Error(w, "Failed to restart deployment", http.StatusBadRequest)
		return
	}

	resp, err := extensions.ReplyForRestart(oldDeployment, newDeployment)
	if err != nil {
		http.Error(w, "Failed to create reply for restart", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("restarted deployment"))

	if err := h.sendMessage(req.Data.Callback.URL, resp); err != nil {
		fmt.Println(err.Error())
	}
}
