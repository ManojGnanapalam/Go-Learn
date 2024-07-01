package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Users struct {
	Name     string `json:"Username"`
	UserId   string
	Email    string
	Password string   `json:"-"`
	Role     []string `json:"Role,omitempty"`
	Address  *Address
}

type Address struct {
	DoorNo int
	City   string
}

var userData []Users

func (c *Users) IsEmpty() bool {
	return c.Name == ""
}

func main() {

}

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Basic CURD APP using GOlang</h1>"))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, user := range userData {
		if user.UserId == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("No user found with given id")
	return

}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Request has no data")
	}

	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	user.UserId = strconv.Itoa(rand.Intn(100))
	userData = append(userData, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, user := range userData {
		if user.UserId == params["id"] {
			userData = append(userData[:i], userData[i+1:]...)
			var user Users
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UserId = params["id"]
			userData = append(userData, user)
			json.NewEncoder(w).Encode(user)
		}
	}
}
