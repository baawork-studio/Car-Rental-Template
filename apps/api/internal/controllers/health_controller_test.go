package controllers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/car-rental-template/api/internal/services"
	"github.com/gin-gonic/gin"
)

type healthCheckerStub struct{ err error }

func (checker healthCheckerStub) Ping(context.Context) error { return checker.err }

func TestGetStatusReturnsHealthResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name, expectedStatus string
		databaseError        error
		wantCode             int
	}{
		{"healthy", "ok", nil, http.StatusOK},
		{"degraded", "degraded", errors.New("database unavailable"), http.StatusServiceUnavailable},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := services.NewHealthService(healthCheckerStub{test.databaseError}, healthCheckerStub{})
			controller := NewHealthController(service)
			router := gin.New()
			router.GET("/health", controller.GetStatus)

			response := httptest.NewRecorder()
			router.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/health", nil))
			if response.Code != test.wantCode {
				t.Fatalf("expected HTTP %d, got %d", test.wantCode, response.Code)
			}
			if !containsJSONStatus(response.Body.String(), test.expectedStatus) {
				t.Fatalf("unexpected response body: %s", response.Body.String())
			}
		})
	}
}

func containsJSONStatus(body, status string) bool {
	return body == `{"status":"`+status+`","database":"ok","cache":"ok"}` || body == `{"status":"`+status+`","database":"unavailable","cache":"ok"}`
}
