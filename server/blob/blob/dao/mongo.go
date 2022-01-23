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
	ativeField = "active"
)

// Mongo defines a mongo dao.
type Mongo struct {
	col *mongo.Collection
}

// NewMongo creates a new mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("blob"),
	}
}

// BlobRecord defines a blob record struct.
type BlobRecord struct {
	mgutil.IDField `bson:"inline"`
	AccountID      id.AccountID `bson:"accountid"`
	PageID         id.PageID    `bson:"pageid"`
	Active         bool         `bson:"active"`
	Path           string       `bson:"path"`
}

// CreateBlob creates a blob record.
func (m *Mongo) CreateBlob(c context.Context, aid id.AccountID, pid id.PageID) (*BlobRecord, error) {
	br := &BlobRecord{
		AccountID: aid,
		PageID:    pid,
	}
	objID := mgutil.NewObjID()
	br.ID = objID
	br.Path = fmt.Sprintf("%s/%s/%s", aid.String(), pid.String(), objID.Hex())

	_, err := m.col.InsertOne(c, br)
	if err != nil {
		return nil, err
	}
	return br, nil
}

// GetBlob gets a blob record.
func (m *Mongo) GetBlob(c context.Context, bid id.BlobID) (*BlobRecord, error) {
	objID, err := objid.FromID(bid)
	if err != nil {
		return nil, fmt.Errorf("invalid object id: %v", err)
	}
	res := m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
	})

	if err := res.Err(); err != nil {
		return nil, err
	}

	var br BlobRecord
	err = res.Decode(&br)
	if err != nil {
		return nil, fmt.Errorf("cannot decode result: %v", err)
	}

	return &br, nil
}

func (m *Mongo) ActiveBlob(c context.Context, bid id.BlobID) error {
	objID, err := objid.FromID(bid)
	if err != nil {
		return fmt.Errorf("invalid object id: %v", err)
	}

	u := bson.M{}
	u[ativeField] = true
	res := m.col.FindOneAndUpdate(c, bson.M{
		mgutil.IDFieldName: objID,
	}, mgutil.Set(u))
	return res.Err()
}
