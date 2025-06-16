package handlers

import (
	"JunieSampleProject/internal/db"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestRouter() (*gin.Engine, *UserHandler) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new router
	router := gin.New()

	// Create a new user store
	userStore := db.NewUserStore()

	// Create a new user handler
	userHandler := NewUserHandler(userStore)

	// Register the routes
	api := router.Group("/api/v1")
	userHandler.RegisterRoutes(api)

	return router, userHandler
}

func TestGetUsers(t *testing.T) {
	router, _ := setupTestRouter()

	// Create a test request
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var users []db.User
	if err := json.Unmarshal(w.Body.Bytes(), &users); err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	// Check that we got the expected number of users
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	router, _ := setupTestRouter()

	// Create a test request
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var user db.User
	if err := json.Unmarshal(w.Body.Bytes(), &user); err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	// Check that we got the expected user
	if user.ID != "1" {
		t.Errorf("Expected user ID 1, got %s", user.ID)
	}
}

func TestCreateUser(t *testing.T) {
	router, _ := setupTestRouter()

	// Create a test request with a JSON body
	newUser := map[string]string{
		"name":  "Test User",
		"email": "test@example.com",
	}
	jsonData, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	// Check the response body
	var createdUser db.User
	if err := json.Unmarshal(w.Body.Bytes(), &createdUser); err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	// Check that the user was created with the correct data
	if createdUser.Name != newUser["name"] || createdUser.Email != newUser["email"] {
		t.Errorf("User data doesn't match: expected %v, got %v", newUser, createdUser)
	}
}

func TestUpdateUser(t *testing.T) {
	router, _ := setupTestRouter()

	// Create a test request with a JSON body
	updateData := map[string]string{
		"name":  "Updated User",
		"email": "updated@example.com",
	}
	jsonData, _ := json.Marshal(updateData)
	req, _ := http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	// Check that the response contains the success message
	if message, ok := response["message"]; !ok || message != "User updated successfully" {
		t.Errorf("Expected success message, got %v", response)
	}
}

func TestDeleteUser(t *testing.T) {
	router, _ := setupTestRouter()

	// Create a test request
	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	// Check that the response contains the success message
	if message, ok := response["message"]; !ok || message != "User deleted successfully" {
		t.Errorf("Expected success message, got %v", response)
	}
}