package dao

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"notion/shared/id"
	mgutil "notion/shared/mongo"
	"notion/shared/mongo/objid"

	accountpb "notion/account/api/gen/v1"
)

const (
	accountField        = "account"
	usernameField       = accountField + ".username"
	emailField          = accountField + ".email"
	passwordField       = accountField + ".password"
	filterPasswordField = "password"
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
	Account        *accountpb.Account `bson:"account"`
}

// CreateAccount creates a account.
func (m *Mongo) CreateAccount(c context.Context, account *AccountRecord) (*AccountRecord, error) {
	account.ID = mgutil.NewObjID()
	_, err := m.col.InsertOne(c, account)
	if err != nil {
		return nil, err
	}
	return account, err
}

// GetAccount gets a account.
func (m *Mongo) GetAccount(c context.Context, id id.AccountID) (*AccountRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	return convertSingleResult(m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
	}))
}

type AccountUpdate struct {
	Username string
	Email    string
}

// UpdateAccount updates a account.
func (m *Mongo) UpdateAccount(c context.Context, id id.AccountID, update *AccountUpdate) (*AccountRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{
		mgutil.IDFieldName: objID,
	}

	u := bson.M{}
	if update.Username != "" {
		u[usernameField] = update.Username
	}
	if update.Email != "" {
		u[emailField] = update.Email
	}

	res := m.col.FindOneAndUpdate(c, filter, mgutil.Set(u),
		options.FindOneAndUpdate().SetReturnDocument(options.After))
	return convertSingleResult(res)
}

// ChangePassword change a account password.
func (m *Mongo) ChangePassword(c context.Context, id id.AccountID, password string) error {
	objID, err := objid.FromID(id)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}

	filter := bson.M{
		mgutil.IDFieldName: objID,
	}

	u := bson.M{}
	u[passwordField] = password
	res := m.col.FindOneAndUpdate(c, filter, mgutil.Set(u))
	return res.Err()
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
