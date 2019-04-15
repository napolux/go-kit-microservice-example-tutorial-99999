package napodate

import (
	"context"
	"encoding/json"
	"net/http"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type getRequest struct{}

type getResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

type validateRequest struct {
	Date string `json:"date"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}

// In the second part we will write "decoders" for our incoming requests
func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	return req, nil
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
