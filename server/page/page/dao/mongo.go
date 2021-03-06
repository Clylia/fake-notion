package dao

import (
	"context"
	"fmt"
	pagepb "notion/page/api/gen/v1"
	"notion/shared/id"
	mgutil "notion/shared/mongo"
	"notion/shared/mongo/objid"

	"github.com/pkg/errors"
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
	mgutil.IDField        `bson:"inline"`
	mgutil.CreatedAtField `bson:"inline"`
	mgutil.UpdatedAtField `bson:"inline"`
	CreatorID             primitive.ObjectID `bson:"creatorid"`
	Blocks                []*BlockEmtity     `bson:"blocks"`
}

// BlockEmtity defines a block emtity.
type BlockEmtity struct {
	mgutil.IDField `bson:"inline"`
	Tag            string `bson:"tag"`
	HTML           string `bson:"html"`
	ImageURL       string `bson:"imageurl"`
}

// CreatePage creates a page.
func (m *Mongo) CreatePage(c context.Context, aid id.AccountID, blocks []*BlockEmtity) (*pagepb.PageEmtity, error) {
	objAID, err := objid.FromID(aid)
	if err != nil {
		return nil, fmt.Errorf("invalid account id[%v]: %w", aid, err)
	}
	var page PageRecord
	page.Blocks = blocks
	page.CreatorID = objAID
	page.ID = mgutil.NewObjID()
	now := mgutil.UpdatedAt()
	page.CreatedAt = now
	page.UpdatedAt = now
	_, err = m.col.InsertOne(c, &page)
	if err != nil {
		return nil, err
	}
	return convertPageEmtity(&page), err
}

// GetPage gets a page.
func (m *Mongo) GetPage(c context.Context, pid id.PageID, aid id.AccountID) (*pagepb.PageEmtity, error) {
	objPID, err := objid.FromID(pid)
	if err != nil {
		return nil, fmt.Errorf("invalid page id[%v]: %w", pid, err)
	}

	objAID, err := objid.FromID(aid)
	if err != nil {
		return nil, fmt.Errorf("invalid account id[%v]: %w", aid, err)
	}
	return convertSingleResult(m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objPID,
		creatorIDField:     objAID,
	}))
}

// GetPages gets pages.
func (m *Mongo) GetPages(c context.Context, aid id.AccountID) ([]*pagepb.PageEmtity, error) {
	objAID, err := objid.FromID(aid)
	if err != nil {
		return nil, fmt.Errorf("invalid account id[%v]: %w", aid, err)
	}
	cur, err := m.col.Find(c, bson.M{
		creatorIDField: objAID,
	})
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("account[%v] cannot find page", aid))
	}

	var pages []*pagepb.PageEmtity
	for cur.Next(c) {
		var page PageRecord
		err = cur.Decode(&page)
		if err != nil {
			return nil, errors.Wrap(err, "cannot decode page")
		}
		pages = append(pages, convertPageEmtity(&page))
	}
	return pages, nil
}

// UpdatePage updates a page.
func (m *Mongo) UpdatePage(c context.Context, pid id.PageID, aid id.AccountID, blocks []*BlockEmtity) (*pagepb.PageEmtity, error) {
	objPID, err := objid.FromID(pid)
	if err != nil {
		return nil, fmt.Errorf("invalid page id[%v]: %w", pid, err)
	}

	objAID, err := objid.FromID(aid)
	if err != nil {
		return nil, fmt.Errorf("invalid account id[%v]: %w", aid, err)
	}
	filter := bson.M{
		mgutil.IDFieldName: objPID,
		creatorIDField:     objAID,
	}

	u := bson.M{}
	if len(blocks) > 0 {
		u[blocksField] = blocks
	}
	u[updateAtField] = mgutil.UpdatedAt()

	res := m.col.FindOneAndUpdate(c, filter, mgutil.Set(u),
		options.FindOneAndUpdate().SetReturnDocument(options.After))
	return convertSingleResult(res)
}

// DeletePage delete a page.
func (m *Mongo) DeletePage(c context.Context, pid id.PageID, aid id.AccountID) error {
	objPID, err := objid.FromID(pid)
	if err != nil {
		return fmt.Errorf("invalid page id[%v]: %w", pid, err)
	}

	objAID, err := objid.FromID(aid)
	if err != nil {
		return fmt.Errorf("invalid account id[%v]: %w", aid, err)
	}

	filter := bson.M{
		mgutil.IDFieldName: objPID,
		creatorIDField:     objAID,
	}

	res := m.col.FindOneAndDelete(c, filter)
	return res.Err()
}

func convertSingleResult(res *mongo.SingleResult) (*pagepb.PageEmtity, error) {
	if err := res.Err(); err != nil {
		return nil, err
	}

	var pr PageRecord
	err := res.Decode(&pr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode: %w", err)
	}

	return convertPageEmtity(&pr), nil
}

func convertPageEmtity(pr *PageRecord) *pagepb.PageEmtity {
	var page pagepb.PageEmtity
	page.Id = pr.ID.Hex()

	for _, b := range pr.Blocks {
		var block pagepb.BlockEmtity
		block.Id = b.ID.Hex()
		block.Html = b.HTML
		block.Tag = b.Tag
		block.ImageUrl = b.ImageURL
		page.Blocks = append(page.Blocks, &block)
	}
	page.CreatedAt = int32(pr.CreatedAt)
	page.UpdatedAt = int32(pr.UpdatedAt)
	return &page
}
