package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

var users = make(map[int]User)
var userIDCounter = 1

func main() {
	http.HandleFunc("/create", createUserHandler)
	http.HandleFunc("/make_friends", makeFriendsHandler)
	http.HandleFunc("/user", deleteUserHandler)
	http.HandleFunc("/friends/", getFriendsHandler)
	http.HandleFunc("/update_age/", updateAgeHandler)
	http.ListenAndServe("localhost:8082", nil)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = userIDCounter
	users[userIDCounter] = user
	userIDCounter++
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"user_id": user.ID})
	w.WriteHeader(http.StatusCreated)
}

func makeFriendsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var requestData struct {
		SourceID int `json:"source_id"`
		TargetID int `json:"target_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if sourceUser, ok := users[requestData.SourceID]; ok {
		if targetUser, ok := users[requestData.TargetID]; ok {
			sourceUser.Friends = append(sourceUser.Friends, requestData.TargetID)
			targetUser.Friends = append(targetUser.Friends, requestData.SourceID)
			users[requestData.SourceID] = sourceUser
			users[requestData.TargetID] = targetUser
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s and %s are now friends.", sourceUser.Name, targetUser.Name)
			return
		}
	}
	http.Error(w, "One or both users not found", http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var requestData struct {
		TargetID int `json:"target_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user, ok := users[requestData.TargetID]; ok {
		delete(users, requestData.TargetID)
		for _, u := range users {
			for i, friendID := range u.Friends {
				if friendID == requestData.TargetID {
					u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				}
			}
			users[u.ID] = u
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Deleted user: %s", user.Name)
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func getFriendsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	userIDStr := r.URL.Path[len("/friends/"):]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	if user, ok := users[userID]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user.Friends)
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func updateAgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	userIDStr := r.URL.Path[len("/update_age/"):]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var requestData struct {
		NewAge int `json:"age"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user, ok := users[userID]; ok {
		user.Age = requestData.NewAge
		users[userID] = user
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User age successfully updated.")
		return
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
