package dataify

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestLocalValidation(t *testing.T) {
	client := NewClient("test-token")

	_, err := client.Google.Flights(context.Background(), GoogleFlightsRequest{})
	if err == nil {
		t.Fatal("expected missing departure_id error")
	}
	if !strings.Contains(err.Error(), "departure_id") {
		t.Fatalf("expected departure_id error, got %v", err)
	}

	_, err = client.Google.Hotels(context.Background(), GoogleHotelsRequest{})
	if err == nil {
		t.Fatal("expected missing check_in_date error")
	}
	if !strings.Contains(err.Error(), "check_in_date") {
		t.Fatalf("expected check_in_date error, got %v", err)
	}

	if !errors.Is(ErrMissingToken, ErrMissingToken) {
		t.Fatal("keep ErrMissingToken referenced for validation build")
	}
}
