package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestCorsBuildsExpectedConfiguration(t *testing.T) {
	config := Cors("https://customer.example,https://admin.example")
	if len(config.AllowOrigins) != 2 || config.AllowOrigins[0] != "https://customer.example" {
		t.Fatalf("unexpected allowed origins: %#v", config.AllowOrigins)
	}
	if len(config.AllowMethods) != 6 || config.AllowMethods[0] != http.MethodGet {
		t.Fatalf("unexpected allowed methods: %#v", config.AllowMethods)
	}
}

func TestCorsHandlesAllowedAndBlockedOrigins(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(cors.New(Cors("https://customer.example")))
	router.GET("/cars", func(context *gin.Context) { context.Status(http.StatusOK) })

	allowed := httptest.NewRecorder()
	allowedRequest := httptest.NewRequest(http.MethodOptions, "/cars", nil)
	allowedRequest.Header.Set("Origin", "https://customer.example")
	allowedRequest.Header.Set("Access-Control-Request-Method", http.MethodGet)
	router.ServeHTTP(allowed, allowedRequest)
	if allowed.Code != http.StatusNoContent || allowed.Header().Get("Access-Control-Allow-Origin") != "https://customer.example" {
		t.Fatalf("expected allowed preflight, got status=%d origin=%q", allowed.Code, allowed.Header().Get("Access-Control-Allow-Origin"))
	}

	blocked := httptest.NewRecorder()
	blockedRequest := httptest.NewRequest(http.MethodOptions, "/cars", nil)
	blockedRequest.Header.Set("Origin", "https://blocked.example")
	blockedRequest.Header.Set("Access-Control-Request-Method", http.MethodGet)
	router.ServeHTTP(blocked, blockedRequest)
	if blocked.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Fatalf("expected blocked origin to have no CORS header, got %q", blocked.Header().Get("Access-Control-Allow-Origin"))
	}
}
