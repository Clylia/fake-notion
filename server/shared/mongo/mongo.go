package mgutil

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"notion/shared/mongo/objid"
)

// Common filed names.
const (
	IDFieldName   = "_id"
	UpdatedAtName = "updatedat"
)

// IDField defines the object id field.
type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

// UpdatedAtField defines the updatedat field.
type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat"`
}

// NewObjID generates a new object id.
var NewObjID = primitive.NewObjectID

// NewObjIDWithValue sets id for next objectID generation.
func NewObjIDWithValue(id fmt.Stringer) {
	NewObjID = func() primitive.ObjectID {
		return objid.MustFromID(id)
	}
}

// UpdatedAt returns a value suitable for UpdatedAt field.
var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

// Set returns a $set update document.
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert returns a $setOnInsert update document.
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

// ZeroOrDoesNotExist generates a filter expression with
// field equal to zero or field does not exists.
func ZeroOrDoesNotExist(field string, zero interface{}) bson.M {
	return bson.M{
		"$or": []bson.M{
			{
				field: zero,
			},
			{
				field: bson.M{
					"$exists": false,
				},
			},
		},
	}
}
