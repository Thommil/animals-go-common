package model

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/thommil/animals-go-common/dao/mongo"
)

// Account model definition
type Account struct {
	// ID of the account
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Provider of the account
	Provider string `json:"provider,omitempty" bson:"provider,omitempty"`
	// ExternalID of the account
	ExternalID string `json:"external_id,omitempty" bson:"external_id,omitempty"`
	// User ID linked to this account
	UserID string `json:"user_id,omitempty" bson:"user_id,omitempty"`
}

// CreateOrUpdateAccount from the parameter and returns the created/updated account
func CreateOrUpdateAccount(account *Account) (*Account, error) {
	var collection = mongo.GetInstance().DB("").C("account")
	var err error
	if account.ID != "" {
		//Update
		err = collection.UpdateId(account.ID, account)
	} else {
		//Create
		account.ID = bson.NewObjectId()
		err = collection.Insert(account)
	}
	if err != nil {
		return nil, err
	}
	return account, nil
}

// FindAccount returns a Mongo Query to parse result, the search query is Mongo based too
func FindAccount(query interface{}) *mgo.Query {
	return mongo.GetInstance().DB("").C("account").Find(query)
}

// FindAccountByID allows to get a single account from its ID
func FindAccountByID(id string) (*Account, error) {
	var collection = mongo.GetInstance().DB("").C("account")
	account := &Account{}
	if bson.IsObjectIdHex(id) {
		err := collection.FindId(bson.ObjectIdHex(id)).One(account)
		if err != nil {
			return nil, err
		}
		return account, nil
	}
	return nil, fmt.Errorf("Invalid account ID")
}

// DeleteAccountByID allows to delete an existing account from its ID
// Deletes also the linked user
func DeleteAccountByID(id string) error {
	var collection = mongo.GetInstance().DB("").C("account")
	if bson.IsObjectIdHex(id) {
		account, err := FindAccountByID(id)
		if err != nil {
			return err
		}
		DeleteUserByID(account.UserID)
		return collection.RemoveId(bson.ObjectIdHex(id))
	}
	return fmt.Errorf("Invalid account ID")
}
