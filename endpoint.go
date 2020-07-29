package basicgokitfirst

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoint struct
type Endpoint struct {
	GetEndpoint      endpoint.Endpoint
	StatusEndpoint   endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getRequest)

		d, err := srv.Get(ctx)

		if err != nil {
			return getResponse{d, err.Error()}, nil
		}

		return getResponse{d, ""}, nil
	}
}

func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest)

		s, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{s}, nil
		}
		return statusResponse{s}, nil
	}
}

func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)

		b, err := srv.Validate(ctx, req.Date)

		if err != nil {
			return validateResponse{b, err.Error()}, nil
		}

		return validateResponse{b, ""}, nil
	}
}

func (e Endpoint) Get(ctx context.Context) (string, error) {
	req := getRequest{}
	resp, err := e.GetEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(getResponse)

	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}

	return getResp.Date, nil
}

func (e Endpoint) Status(ctx context.Context) (string, error) {
	req := statusRequest{}
	resp, err := e.StatusEndpoint(ctx, req)

	if err != nil {
		return "", err
	}

	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

func (e Endpoint) Validate(ctx context.Context, date string) (bool, error) {
	req := validateRequest{Date: date}

	resp, err := e.ValidateEndpoint(ctx, req)

	if err != nil {
		return false, err
	}

	validateResp := resp.(validateResponse)

	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}

	return validateResp.Valid, nil
}
