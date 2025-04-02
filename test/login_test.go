package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Happy version
func TestLogin(t *testing.T) {
	router := SetUpRoutes()
	// set up login payload
	login_payload := []byte(`{
		"phone":"0906371703",
		"password":"HJ10xugb123*"
	}`)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(login_payload))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

// Wrong password
func TestLoginWrongPassword(t *testing.T) {
	router := SetUpRoutes()
	// set up payload
	login_payload := []byte(`{
		"phone":"0906371703",
		"password":"HJ116112002xugb123*"
	}`)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(login_payload))
	if err != nil {
		t.Fatalf("Could not create or send request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Record the response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	var body map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("Faile to decode reponse: %v", err)
	}
	if body["error"] == nil {
		t.Error("Expected an error message in response")
	} else {
		t.Logf("Error Message %v", body["error"])
	}
}

// Invalid Phone number
func TestLoginInvalidPhone(t *testing.T) {
	routers := SetUpRoutes()
	// set up payload
	login_payload := []byte(`{
		"phone":"12343567899",
		"password":"HJ10xugb123*"
	}`)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(login_payload))
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	//Get the record
	rec := httptest.NewRecorder()
	routers.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad request, got %d", rec.Code)
	}
	var body map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	} else {
		t.Logf("Error Message %v", body["error"])
	}
}

// Empty Phone number
func TestEmptyField(t *testing.T) {
	routers := SetUpRoutes()
	// Prepare payload
	login_payload := []byte(`{
		"phone":"",
		"password" :""
	}`)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(login_payload))
	if err != nil {
		t.Fatalf("Can not send request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	routers.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %d", rec.Code)
	}
	var body map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	} else {
		t.Logf("Error Message: %v", body["error"])
	}
}
