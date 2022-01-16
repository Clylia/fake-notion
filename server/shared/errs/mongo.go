package errs

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// IsNoDocuments checkes error is mongo.ErrNoDocuments error.
func IsNoDocumentsErr(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}

// IsDuplicateKeyErr checkes error is duplicate key error.
func IsDuplicateKeyErr(err error) bool {
	return mongo.IsDuplicateKeyError(err)
}
