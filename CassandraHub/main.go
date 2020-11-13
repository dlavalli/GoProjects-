package main

/*
GET / 				-- heartbeat check if our API is online GET /users -- fetch all users from the database
GET /users/UUID 	-- fetch an individual user from the database
POST /users/new 	-- create a new user GET /messages -- fetch all messages from Stream (with the database as a backup)
GET /messages/UUID 	-- fetch an individual message from the database POST /messages/new -- create a new message
*/

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/MyGoProjects/CassandraHub/Cassandra"
	"github.com/MyGoProjects/CassandraHub/Users"
	"log"
	"net/http"
)

// In order for gorilla/mux to return any payload, we must define the exact structure of the data.
// The <code> json:”status” annotation portion of the field definition above will tell our JSON-encoder
// to rename the “Status” and “Code” fields during the encoding process to their lowercase versions.
type heartbeatResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func main() {
	// register to close the session later on when the function enclosing the defer is completed
	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	// set up a new gorilla/mux router to handle our API endpoints according to the contract we described above.
	// We can add a new API route after our router is declared, and give it the name of another function to execute
	// when a request is made for that endpoint.
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)

	// Start listening on the specified port or exist with an error log
	log.Fatal(http.ListenAndServe(":8080", router))
}


