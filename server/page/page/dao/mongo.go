package dao

import (
	"context"
	"fmt"
	pagepb "notion/page/api/gen/v1"
	"notion/shared/id"
	mgutil "notion/shared/mongo"
	"notion/shared/mongo/objid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	blocksField    = "blocks"
	updateAtField  = "updateat"
	creatorIDField = "creatorid"
)

// Mongo defines a mongo dao.
type Mongo struct {
	col *mongo.Collection
}

// NewMongo creates a new mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("page"),
	}
}

// PageRecord defines a page record in mongo db.
type PageRecord struct {
	mgutil.IDField `bson:"inline"`
	CreatorID      primitive.ObjectID    `bson:"creatorid"`
	Blocks         []*pagepb.BlockEmtity `bson:"blocks"`
	CreatedAt      int64                 `bson:"createdat"`
	UpdatedAt      int64                 `bson:"updatedat"`
}

// CreatePage creates a page.
func (m *Mongo) CreatePage(c context.Context, page *PageRecord) (*PageRecord, error) {
	page.ID = mgutil.NewObjID()
	_, err := m.col.InsertOne(c, page)
	if err != nil {
		return nil, err
	}
	return page, err
}

// GetPage gets a page.
func (m *Mongo) GetPage(c context.Context, id id.PageID) (*PageRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid page id[%v]: %w", id, err)
	}
	return convertSingleResult(m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
	}))
}

type PageUpdate struct {
	Blocks    []*pagepb.BlockEmtity
	CreatorID id.AccountID
}

// UpdatePage updates a page.
func (m *Mongo) UpdatePage(c context.Context, id id.PageID, update *PageUpdate) (*PageRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid page id[%v]: %w", id, err)
	}

	objAID, err := objid.FromID(update.CreatorID)
	if err != nil {
		return nil, fmt.Errorf("invalid account id[%v]: %w", id, err)
	}
	filter := bson.M{
		mgutil.IDFieldName: objID,
		creatorIDField:     objAID,
	}

	u := bson.M{}
	if len(update.Blocks) > 0 {
		u[blocksField] = update.Blocks
	}
	u[updateAtField] = mgutil.UpdatedAt()

	res := m.col.FindOneAndUpdate(c, filter, mgutil.Set(u),
		options.FindOneAndUpdate().SetReturnDocument(options.After))
	return convertSingleResult(res)
}

// DeletePage delete a page.
func (m *Mongo) DeletePage(c context.Context, pid id.PageID, aid id.AccountID) error {
	pageObjID, err := objid.FromID(pid)
	if err != nil {
		return fmt.Errorf("invalid page id[%v]: %w", pid, err)
	}

	accountObjID, err := objid.FromID(aid)
	if err != nil {
		return fmt.Errorf("invalid account id[%v]: %w", aid, err)
	}

	filter := bson.M{
		mgutil.IDFieldName: pageObjID,
		creatorIDField:     accountObjID,
	}

	res := m.col.FindOneAndDelete(c, filter)
	return res.Err()
}

func convertSingleResult(res *mongo.SingleResult) (*PageRecord, error) {
	if err := res.Err(); err != nil {
		return nil, err
	}

	var ar PageRecord
	err := res.Decode(&ar)
	if err != nil {
		return nil, fmt.Errorf("cannot decode: %w", err)
	}
	return &ar, nil
}
