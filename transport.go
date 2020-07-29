package basicgokitfirst

import (
	"context"
	"encoding/json"
	"net/http"
)

// first part define/mapping request and response for each function
type getRequest struct{}

type getResponse struct {
	Date string `json: "date"`
	Err  string `json: "err, omitempty"`
}

type validateRequest struct {
	Date string `json: "date"`
}

type validateResponse struct {
	Valid bool   `json: "valid"`
	Err   string `json: "err, omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json: "status"`
}

//second part for write "docoders" from incoming request
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

// third step for encode response
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
