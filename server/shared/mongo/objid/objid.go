package objid

import (
	"fmt"
	"notion/shared/id"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FromID converts an id to objected id.
func FromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

// MustFromID converts an id to objected id, panics on error.
func MustFromID(id fmt.Stringer) primitive.ObjectID {
	oid, err := FromID(id)
	if err != nil {
		panic(err)
	}

	return oid
}

// ToAccountID converts object id to account id.
func ToAccountID(oid primitive.ObjectID) id.AccountID {
	return id.AccountID(oid.Hex())
}

// ToPageID converts object id to page id.
func ToPageID(oid primitive.ObjectID) id.PageID {
	return id.PageID(oid.Hex())
}
