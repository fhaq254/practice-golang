// https://tutorialedge.net/golang/golang-orm-tutorial/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserList []User

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tutorial - Mux")
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	user := UserList{
		User{Name: "jame", Email: "jame@gmail.com"},
		User{Name: "john", Email: "john@gmail.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User is Created successfully.")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User is Updated successfully.")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User is Deleted successfully.")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home).Methods("GET")
	myRouter.HandleFunc("/user", getUserList).Methods("GET")
	myRouter.HandleFunc("/user/{name}", createUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
