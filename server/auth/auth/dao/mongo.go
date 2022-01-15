package dao

import (
	"context"
	"fmt"
	"notion/shared/id"
	mgutil "notion/shared/mongo"
	"notion/shared/mongo/objid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	accountField = "account"
	emailField   = accountField + ".email"
)

// Mongo defines a mongo dao.
type Mongo struct {
	col *mongo.Collection
}

// NewMongo creates a new mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

// AccountRecord defines a account record in mongo db.
type AccountRecord struct {
	mgutil.IDField `bson:"inline"`
	Account        struct {
		Username string `bson:"username"`
		Email    string `bson:"email"`
		Password string `bson:"password"`
	} `bson:"account"`
}

// ResolveAccount resolves an account id from email.
func (m *Mongo) ResolveAccount(c context.Context, email string) (*AccountRecord, error) {
	return convertSingleResult(m.col.FindOne(c, bson.M{
		emailField: email,
	}))
}

// Exists return account record if exists.
func (m *Mongo) Exists(c context.Context, id id.AccountID) (*AccountRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	return convertSingleResult(m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
	}))
}

func convertSingleResult(res *mongo.SingleResult) (*AccountRecord, error) {
	if err := res.Err(); err != nil {
		return nil, err
	}

	var ar AccountRecord
	err := res.Decode(&ar)
	if err != nil {
		return nil, fmt.Errorf("cannot decode: %w", err)
	}
	return &ar, nil
}
