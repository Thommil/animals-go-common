package model

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User model definition
type User struct {
	// ID of the user
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Username of the user
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	// Picture URL of the user
	Picture string `json:"picture,omitempty" bson:"picture,omitempty"`
	// Locale of the user
	Locale string `json:"locale,omitempty" bson:"locale,omitempty"`
}

// CreateOrUpdateUser from the parameter and returns the created/updated user
func CreateOrUpdateUser(database *mgo.Database, user *User) (*User, error) {
	var err error
	if user.ID != "" {
		//Update
		err = database.C("user").UpdateId(user.ID, user)
	} else {
		//Create
		user.ID = bson.NewObjectId()
		err = database.C("user").Insert(user)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUser returns a Mongo Query to parse result, the search query is Mongo based too
func FindUser(database *mgo.Database, query interface{}) *mgo.Query {
	return database.C("user").Find(query)
}

// FindUserByID allows to get a single user from her ID
func FindUserByID(database *mgo.Database, id string) (*User, error) {
	user := &User{}
	if bson.IsObjectIdHex(id) {
		err := database.C("user").FindId(bson.ObjectIdHex(id)).One(user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, fmt.Errorf("Invalid user ID")
}

// DeleteUserByID allows to delete an existing user from her ID
func DeleteUserByID(database *mgo.Database, id string) error {
	if bson.IsObjectIdHex(id) {
		return database.C("user").RemoveId(bson.ObjectIdHex(id))
	}
	return fmt.Errorf("Invalid user ID")
}
