package services

import (
	"context"
	"errors"
	"testing"
)

type stubHealthChecker struct{ err error }

func (checker stubHealthChecker) Ping(context.Context) error { return checker.err }

func TestHealthServiceStatus(t *testing.T) {
	tests := []struct {
		name                                string
		postgresError, redisError           error
		wantStatus, wantDatabase, wantCache string
	}{
		{"all services available", nil, nil, "ok", "ok", "ok"},
		{"database unavailable", errors.New("database down"), nil, "degraded", "unavailable", "ok"},
		{"cache unavailable", nil, errors.New("cache down"), "degraded", "ok", "unavailable"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewHealthService(stubHealthChecker{test.postgresError}, stubHealthChecker{test.redisError})
			status := service.Status(context.Background())
			if status.Status != test.wantStatus || status.Database != test.wantDatabase || status.Cache != test.wantCache {
				t.Fatalf("unexpected status: %#v", status)
			}
		})
	}
}
