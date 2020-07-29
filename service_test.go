package basicgokitfirst

import (
	"context"
	"testing"
)

func TestStatus(t *testing.T) {

	// srv, ctx := Setup()
	var resultContext context.Context = context.Background()
	dateService := DateService{}
	result, err := dateService.Status(resultContext)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	success := result == "ok"
	if !success {
		t.Errorf("expected service to be ok!")
	}
}

// func Setup() (srv Service, ctx context.Context) {
// 	return NewService(), context.Background()
// }
