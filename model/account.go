package model

import "github.com/globalsign/mgo/bson"

// Account model definition
type Account struct {
	// ID of the account
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Provider of the account
	Provider string `json:"provider,omitempty" bson:"provider,omitempty"`
	// ExternalID of the account
	ExternalID string `json:"external_id,omitempty" bson:"external_id,omitempty"`
	// User ID linked to this account
	UserID string `json:"user_id" bson:"user_id"`
}
