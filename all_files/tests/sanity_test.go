package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"taskmanage/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func init() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}


func TestEndpointAvailability(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loaffffing .env file: %v", err)
	}
	router := gin.Default()
	routes.SetupRoutes(router)
	endpoints := []string{"/login", "/register", "/api/status"}
	for _, endpoint := range endpoints {
		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			t.Fatalf("Failed to create request for %s: %v", endpoint, err)
		}

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200 for %s, got %d", endpoint, rec.Code)
		}
	}
}
