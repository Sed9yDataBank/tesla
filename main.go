package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/faygun/go-rest-api/helper"
	"github.com/faygun/go-rest-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDB()

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["userid"])
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["userid"])
	var user models.User
	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)
	update := bson.D{
		{"$set", bson.D{
			{"teslamodel", user.TeslaModel},
			{"location", user.Location},
		}},
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	user.UserID = id
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["userid"])
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}

func enableUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["userid"])
	var user models.User
	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)
	update := bson.D{
		{"$set", bson.D{
			{"enabled", user.Enabled},
		}},
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		helper.GetError(err, w)
		return
	}
	user.UserID = id
	json.NewEncoder(w).Encode(user)
}

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{userid}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{userid}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{userid}", deleteUser).Methods("DELETE")
	r.HandleFunc("/api/users/enable/{userid}", enableUser).Methods("PUT")

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
