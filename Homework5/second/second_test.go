package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	reqBody := bytes.NewBuffer([]byte(`{"name": "Test User", "age": 30}`))
	req := httptest.NewRequest("POST", "/create", reqBody)
	recorder := httptest.NewRecorder()

	createUserHandler(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}

	var responseData map[string]int
	err := json.NewDecoder(recorder.Body).Decode(&responseData)
	if err != nil {
		t.Error(err)
	}

	if responseData["user_id"] < 1 {
		t.Error("Expected a positive user ID, got an invalid value")
	}
}

func TestMakeFriendsHandler(t *testing.T) {
	// Create two users for testing
	createUserHandler(nil, httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte(`{"name": "Test User 1", "age": 30}`))))
	createUserHandler(nil, httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte(`{"name": "Test User 2", "age": 25}`))))

	reqBody := bytes.NewBuffer([]byte(`{"source_id": 1, "target_id": 2}`))
	req := httptest.NewRequest("POST", "/make_friends", reqBody)
	recorder := httptest.NewRecorder()

	makeFriendsHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestDeleteUserHandler(t *testing.T) {
	// Create a user for testing
	createUserHandler(nil, httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte(`{"name": "Test User", "age": 30}`))))

	reqBody := bytes.NewBuffer([]byte(`{"target_id": 1}`))
	req := httptest.NewRequest("DELETE", "/user", reqBody)
	recorder := httptest.NewRecorder()

	deleteUserHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestGetFriendsHandler(t *testing.T) {
	// Create two users for testing
	createUserHandler(nil, httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte(`{"name": "Test User 1", "age": 30}`))))
	createUserHandler(nil, httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte(`{"name": "Test User 2", "age": 25}`))))

	// Make them friends
	makeFriendsHandler(nil, httptest.NewRequest("POST", "/make_friends", bytes.NewBuffer([]byte(`{"source_id": 1, "target_id": 2}`))))

	req := httptest.NewRequest("GET", "/friends/1", nil)
	recorder := httptest.NewRecorder()

	getFriendsHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestUpdateAgeHandler(t *testing.T) {
	// Create a user for testing
	createUserHandler(nil, httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte(`{"name": "Test User", "age": 30}`))))

	reqBody := bytes.NewBuffer([]byte(`{"age": 35}`))
	req := httptest.NewRequest("PUT", "/update_age/1", reqBody)
	recorder := httptest.NewRecorder()

	updateAgeHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
