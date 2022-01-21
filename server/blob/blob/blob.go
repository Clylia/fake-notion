package blob

import (
	"context"
	blobpb "notion/blob/api/gen/v1"

	"go.uber.org/zap"
)

// Service defines a account service.
type Service struct {
	blobpb.UnimplementedBlobServiceServer
	Logger *zap.Logger
}

// CreateBlob creates a blob.
func (s *Service) CreateBlob(c context.Context, req *blobpb.CreateBlobRequest) (*blobpb.CreateBlobResponse, error) {
	panic("not implemented") // TODO: Implement
}

// DeleteBlob deletes a blob.
func (s *Service) DeleteBlob(c context.Context, req *blobpb.DeleteBlobRequest) (*blobpb.DeleteBlobResponse, error) {
	panic("not implemented") // TODO: Implement
}
