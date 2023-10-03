package mongodb

import (
	"github.com/tot0p/env"
	"testing"
)

func TestMongoDBPing(t *testing.T) {
	err := env.LoadPath("../../../.env")
	if err != nil {
		t.Error(err)
	}

	m, err := NewMongoDB(env.Get("URI_MONGODB"))
	if err != nil {
		t.Error(err)
	}
	err = m.Ping()
	if err != nil {
		t.Error(err)
	}
	err = m.Disconnect()
	if err != nil {
		t.Error(err)
	}
}
