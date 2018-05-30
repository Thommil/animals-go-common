package model

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
func CreateOrUpdateAccount(database *mgo.Database, account *Account) (*Account, error) {
	var err error
	if account.ID != "" {
		//Update
		err = database.C("account").UpdateId(account.ID, account)
	} else {
		//Create
		account.ID = bson.NewObjectId()
		err = database.C("account").Insert(account)
	}
	if err != nil {
		return nil, err
	}
	return account, nil
}

// FindAccount returns a Mongo Query to parse result, the search query is Mongo based too
func FindAccount(database *mgo.Database, query interface{}) *mgo.Query {
	return database.C("account").Find(query)
}

// FindAccountByID allows to get a single account from its ID
func FindAccountByID(database *mgo.Database, id string) (*Account, error) {
	account := &Account{}
	if bson.IsObjectIdHex(id) {
		err := database.C("account").FindId(bson.ObjectIdHex(id)).One(account)
		if err != nil {
			return nil, err
		}
		return account, nil
	}
	return nil, fmt.Errorf("Invalid account ID")
}

// DeleteAccountByID allows to delete an existing account from its ID
// Deletes also the linked user
func DeleteAccountByID(database *mgo.Database, id string) error {
	if bson.IsObjectIdHex(id) {
		account, err := FindAccountByID(database, id)
		if err != nil {
			return err
		}
		DeleteUserByID(database, account.UserID)
		return database.C("account").RemoveId(bson.ObjectIdHex(id))
	}
	return fmt.Errorf("Invalid account ID")
}
