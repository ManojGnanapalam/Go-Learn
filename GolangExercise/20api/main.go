package main

import (
	"encoding/json"
	"log"
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

func (u *Users) IsEmpty() bool {
	return u.Name == ""
}

// func (u *Users) IsIdEqueal(Id string) bool{

// }

func main() {
	r := mux.NewRouter()

	userData = append(userData, Users{UserId: "2", Name: "viper", Email: "viper@goo.co", Role: []string{"user", "editer"}, Address: &Address{DoorNo: 14, City: "chennai"}})

	r.HandleFunc("/", MainPage).Methods("GET")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getOneUser).Methods("GET")
	r.HandleFunc("/user", createNewUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteOneUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
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
		}
	}
	json.NewEncoder(w).Encode("No user found with given id")

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
	isIdFound := false
	params := mux.Vars(r)

	if r.Body == nil {
		json.NewEncoder(w).Encode("Request has no data")
	}

	for i, user := range userData {
		if user.UserId == params["id"] {
			isIdFound = true
			userData = append(userData[:i], userData[i+1:]...)

			var user Users
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UserId = params["id"]
			userData = append(userData, user)
			json.NewEncoder(w).Encode(user)
		}
	}

	if !isIdFound {
		json.NewEncoder(w).Encode("Id is not founded")
	}
}

func deleteOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, user := range userData {
		if user.UserId == params["id"] {
			userData = append(userData[:i], userData[i+1:]...)
			break
		}
	}
}
