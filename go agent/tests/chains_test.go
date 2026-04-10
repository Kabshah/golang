package tests

import (
	"testing"

	"github.com/google/uuid"

	"go-agent/chains"
)

// TestGetVacationFromDb_Found tests retrieving an existing vacation
func TestGetVacationFromDb_Found(t *testing.T) {
	// Clear global slice and add a test vacation
	chains.Vacations = []*chains.Vacation{}
	testID := uuid.New()
	testVacation := &chains.Vacation{Id: testID, Completed: true, Idea: "Test idea"}
	chains.Vacations = append(chains.Vacations, testVacation)

	// Retrieve the vacation
	result, err := chains.GetVacationFromDb(testID)
	if err != nil {
		t.Errorf("Expected no error getting vacation, got %v", err)
	}
	if result.Id != testID {
		t.Errorf("Expected ID %v, got %v", testID, result.Id)
	}
	if result.Completed != true {
		t.Errorf("Expected Completed to be true, got %v", result.Completed)
	}
	if result.Idea != "Test idea" {
		t.Errorf("Expected Idea to be 'Test idea', got '%s'", result.Idea)
	}
}

// TestGetVacationFromDb_NotFound tests retrieving a non-existent vacation
func TestGetVacationFromDb_NotFound(t *testing.T) {
	// Clear global slice
	chains.Vacations = []*chains.Vacation{}
	testID := uuid.New()

	// Try to retrieve non-existent vacation
	_, err := chains.GetVacationFromDb(testID)
	if err == nil {
		t.Error("Expected error for non-existent ID, got nil")
	}
	if err.Error() != "ID not found" {
		t.Errorf("Expected error 'ID not found', got '%v'", err)
	}
}

// Test chains structs
func TestVacation_Fields(t *testing.T) {
	id := uuid.New()
	v := chains.Vacation{
		Id:        id,
		Completed: true,
		Idea:      "Test vacation idea",
	}

	if v.Id != id {
		t.Errorf("Expected Id to be %v, got %v", id, v.Id)
	}
	if v.Completed != true {
		t.Errorf("Expected Completed to be true, got %v", v.Completed)
	}
	if v.Idea != "Test vacation idea" {
		t.Errorf("Expected Idea to be 'Test vacation idea', got '%s'", v.Idea)
	}
}