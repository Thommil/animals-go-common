package model

import (
	"fmt"

	"github.com/thommil/animals-go-common/dao/mongo"

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
func CreateOrUpdateUser(user *User) (*User, error) {
	var collection = mongo.GetInstance().DB("").C("user")
	var err error
	if user.ID != "" {
		//Update
		err = collection.UpdateId(user.ID, user)
	} else {
		//Create
		user.ID = bson.NewObjectId()
		err = collection.Insert(user)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUser returns a Mongo Query to parse result, the search query is Mongo based too
func FindUser(query interface{}) *mgo.Query {
	return mongo.GetInstance().DB("").C("user").Find(query)
}

// FindUserByID allows to get a single user from her ID
func FindUserByID(id string) (*User, error) {
	var collection = mongo.GetInstance().DB("").C("user")
	user := &User{}
	if bson.IsObjectIdHex(id) {
		err := collection.FindId(bson.ObjectIdHex(id)).One(user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, fmt.Errorf("Invalid user ID")
}

// DeleteUserByID allows to delete an existing user from her ID
func DeleteUserByID(id string) error {
	var collection = mongo.GetInstance().DB("").C("user")
	if bson.IsObjectIdHex(id) {
		return collection.RemoveId(bson.ObjectIdHex(id))
	}
	return fmt.Errorf("Invalid user ID")
}
