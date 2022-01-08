// https://tutorialedge.net/golang/golang-orm-tutorial/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database !")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tutorial - Gorm & Mux")
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database !")
	}
	defer db.Close()

	var user []User
	db.Find(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database !")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	email := fmt.Sprintf("%v@gmail.com", name)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User is Created successfully.")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database !")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)

	email := fmt.Sprintf("%v@coolest-company.com", name)
	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "User is Updated successfully.")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database !")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

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
	initialMigration()
	handleRequest()
}
