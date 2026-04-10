package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"go-agent/routes"
)

func TestGenerateVacationHandler(t *testing.T) {
	// Set up gin
	gin.SetMode(gin.TestMode)
	router := gin.New()
	routes.GetVacationRouter(router)

	// Create request body
	body := `{
		"favourite_season": "summer",
		"hobbies": ["hiking", "swimming"],
		"budget": 2000
	}`

	req, _ := http.NewRequest(http.MethodPost, "/vacation/create", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK (%d), got %d", http.StatusOK, w.Code)
	}

	// Parse response JSON (we could unmarshal, but for simplicity just check it's not empty)
	if w.Body.Len() == 0 {
		t.Error("Expected non-empty response body")
	}
}

func TestGetVacationHandler_NotFound(t *testing.T) {
	// Set up gin
	gin.SetMode(gin.TestMode)
	router := gin.New()
	routes.GetVacationRouter(router)

	// Create request with non-existent ID
	req, _ := http.NewRequest(http.MethodGet, "/vacation/00000000-0000-0000-0000-000000000000", nil)

	// Create response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check response - should be not found
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status NotFound (%d), got %d", http.StatusNotFound, w.Code)
	}
}

func TestGetVacationHandler_BadRequest(t *testing.T) {
	// Set up gin
	gin.SetMode(gin.TestMode)
	router := gin.New()
	routes.GetVacationRouter(router)

	// Create request with invalid UUID
	req, _ := http.NewRequest(http.MethodGet, "/vacation/invalid-id", nil)

	// Create response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check response - should be bad request
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest (%d), got %d", http.StatusBadRequest, w.Code)
	}
}

// Test routes structs
func TestGenerateVacationIdeaRequest_Fields(t *testing.T) {
	req := routes.GenerateVacationIdeaRequest{
		FavoriteSeason: "summer",
		Hobbies:        []string{"hiking", "swimming"},
		Budget:         2000,
	}

	if req.FavoriteSeason != "summer" {
		t.Errorf("Expected FavoriteSeason to be 'summer', got '%s'", req.FavoriteSeason)
	}
	if len(req.Hobbies) != 2 {
		t.Errorf("Expected Hobbies length to be 2, got %d", len(req.Hobbies))
	}
	if req.Hobbies[0] != "hiking" {
		t.Errorf("Expected first hobby to be 'hiking', got '%s'", req.Hobbies[0])
	}
	if req.Budget != 2000 {
		t.Errorf("Expected Budget to be 2000, got %d", req.Budget)
	}
}

func TestGenerateVacationIdeaResponse_Fields(t *testing.T) {
	id := uuid.New()
	resp := routes.GenerateVacationIdeaResponse{
		Id:        id,
		Completed: false,
	}

	if resp.Id != id {
		t.Errorf("Expected Id to be %v, got %v", id, resp.Id)
	}
	if resp.Completed != false {
		t.Errorf("Expected Completed to be false, got %v", resp.Completed)
	}
}

func TestGetVacationIdeaResponse_Fields(t *testing.T) {
	id := uuid.New()
	resp := routes.GetVacationIdeaResponse{
		Id:        id,
		Completed: true,
		Idea:      "A great vacation idea",
	}

	if resp.Id != id {
		t.Errorf("Expected Id to be %v, got %v", id, resp.Id)
	}
	if resp.Completed != true {
		t.Errorf("Expected Completed to be true, got %v", resp.Completed)
	}
	if resp.Idea != "A great vacation idea" {
		t.Errorf("Expected Idea to be 'A great vacation idea', got '%s'", resp.Idea)
	}
}