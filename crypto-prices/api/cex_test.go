package api_test

import (
	"cryptomasters/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Expected error for empty currency")
	}
}