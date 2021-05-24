package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	// create multiplexer from mux
	multiplexer := mux.NewRouter()
	// route for index
	multiplexer.HandleFunc("/", Index).Methods("GET")
	// routes for api
	// rest api for users
	multiplexer.HandleFunc("/api/v1/users", GetUsers).Methods("GET")
	multiplexer.HandleFunc("/api/v1/users/{id}", GetUser).Methods("GET")
	multiplexer.HandleFunc("/api/v1/users/{id}", CreateUser).Methods("POST")
	multiplexer.HandleFunc("/api/v1/users/{id}", UpdateUser).Methods("PUT")
	multiplexer.HandleFunc("/api/v1/users/{id}", DeleteUser).Methods("DELETE")

	http.ListenAndServe("127.0.0.1:8000", multiplexer)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: CreateUser page")
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, u := range Users {
		if u.Id == id {
			json.NewEncoder(w).Encode(map[string]string{"message": "user have been created"})
			return
		}
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	user.Id = id
	Users = append(Users, user)
	json.NewEncoder(w).Encode(map[string]string{"message": "user have created"})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: GetUsers page")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: GetUser page")
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for _, user := range Users {
		if user.Id == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "No such user or invalid data"})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: UpdateUser page")
	w.Header().Add("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	id := mux.Vars(r)["id"]
	var temp User
	json.Unmarshal(reqBody, &temp)
	for index, user := range Users {
		if user.Id == id {
			temp.Id = id
			Users[index] = temp
			json.NewEncoder(w).Encode(map[string]string{"message": "user have updated"})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "No such user or invalid data"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: DeleteUser page")
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "user have deleted"})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "No such user or invalid data"})
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[#] Endpoint hit: Index page")
	fmt.Fprintf(w, "Index page")
}
