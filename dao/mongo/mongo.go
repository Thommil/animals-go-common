package mongo

import (
	"sync"

	"github.com/globalsign/mgo"
)

// Configuration definition for MongoDB DAO
type Configuration struct {
	URL string
}

var instance *mgo.Session
var once sync.Once

// NewInstance initialize Mongo instance and returns it
func NewInstance(configuration *Configuration) (*mgo.Session, error) {
	var err error
	once.Do(func() {
		instance, err = mgo.Dial(configuration.URL)
	})
	return instance, err
}

// GetInstance returns the Mongo instance
func GetInstance() *mgo.Session {
	return instance
}
