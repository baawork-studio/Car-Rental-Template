package routes

import (
	"testing"

	"github.com/car-rental-template/api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func TestRegisterAddsHealthRoutes(t *testing.T) {
	router := gin.New()
	Register(router, &controllers.HealthController{})

	expectedPaths := map[string]bool{
		"GET /health":        false,
		"GET /api/v1/health": false,
	}
	for _, route := range router.Routes() {
		key := route.Method + " " + route.Path
		if _, ok := expectedPaths[key]; ok {
			expectedPaths[key] = true
		}
	}
	for route, registered := range expectedPaths {
		if !registered {
			t.Errorf("expected route %s to be registered", route)
		}
	}
}
