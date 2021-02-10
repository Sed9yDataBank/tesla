package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	PostID          primitive.ObjectID `json:"postid,omitempty" bson:"postid,omitempty"`
	PostName        string             `json:"postname" bson:"postname,omitempty"`
	PostUrl         string             `json:"posturl" bson:"posturl,omitempty"`
	PostDescription string             `json:"postdescription" bson:"postdescription,omitempty"`
	User            *User              `json:"user" bson:"user,omitempty"`
	Subreddit       *Subreddit         `json:"subreddit" bson:"subreddit,omitempty"`
}

type Comment struct {
	CommentID primitive.ObjectID `json:"commentid,omitempty" bson:"commentid,omitempty"`
	Text      string             `json:"text" bson:"text,omitempty"`
	Post      *Post              `json:"post" bson:"post,omitempty"`
	User      *User              `json:"user" bson:"user,omitempty"`
}

type User struct {
	UserID     primitive.ObjectID `json:"userid,omitempty" bson:"userid,omitempty"`
	Username   string             `json:"username" bson:"username,omitempty"`
	Password   string             `json:"password" bson:"password,omitempty"`
	Email      string             `json:"email" bson:"email,omitempty"`
	TeslaModel string             `json:"teslamodel" bson:"teslamodel,omitempty"`
	Location   string             `json:"location" bson:"location,omitempty"`
	//Enabled is nil
	Enabled *bool `json:"enabled" bson:"enabled,omitempty"`
}

type Subreddit struct {
	SubredditID primitive.ObjectID `json:"subredditid,omitempty" bson:"subredditid,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Posts       []*Post            `json:"posts" bson:"posts,omitempty"`
	User        *User              `json:"user" bson:"user,omitempty"`
}
