package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Connection mongoDB with helper class
collection := helper.ConnectDB() 

func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{userid}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{userid}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{userid}", deleteUser).Methods("DELETE")

	r.HandleFunc("/api/comments/{postid}", getComments).Methods("GET")
	r.HandleFunc("/api/comments/{commentid}", getComment).Methods("GET")
	r.HandleFunc("/api/comments/{postid}", createComment).Methods("POST")
	r.HandleFunc("/api/comments/{commentid}", updateComment).Methods("PUT")
	r.HandleFunc("/api/comments/{commentid}", deleteComment).Methods("DELETE")

	r.HandleFunc("/api/posts", getPosts).Methods("GET")
	r.HandleFunc("/api/posts/{postid}", getPost).Methods("GET")
	r.HandleFunc("/api/posts/{subredditid}", createPost).Methods("POST")
	r.HandleFunc("/api/posts/{postid}", updatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{postid}", deletePost).Methods("DELETE")

	r.HandleFunc("/api/subreddits", getSubreddits).Methods("GET")
	r.HandleFunc("/api/subreddits/{subredditid}", getSubreddit).Methods("GET")
	r.HandleFunc("/api/subreddits", createSubreddit).Methods("POST")
	r.HandleFunc("/api/subreddits/{subredditid}", updateSubreddit).Methods("PUT")
	r.HandleFunc("/api/subreddits/{subredditid}", deleteSubreddit).Methods("DELETE")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}