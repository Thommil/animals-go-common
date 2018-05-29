package model

import "github.com/globalsign/mgo/bson"

// User model definition
type User struct {
	// ID of the user
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Username of the user
	Username string `json:"username" bson:"username"`
}
