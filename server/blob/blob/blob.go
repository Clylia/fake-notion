package blob

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	blobpb "notion/blob/api/gen/v1"
	"notion/shared/auth"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Storage defines storage interface.
type Storage interface {
	SignURL(c context.Context, method, path string, timeout time.Duration) (string, error)
	Get(c context.Context, path string) (io.ReadCloser, error)
}

// Service defines a blob service
type Service struct {
	blobpb.UnimplementedBlobServiceServer
	Storage Storage
	Logger  *zap.Logger
}

// CreateBlob creates a blob.
func (s *Service) CreateBlob(c context.Context, req *blobpb.CreateBlobRequest) (*blobpb.CreateBlobResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}
	path := fmt.Sprintf("%s/%s/%s", aid, req.PageId, uid)
	u, err := s.Storage.SignURL(c, http.MethodPut, path, secToDuration(req.UploadUrlTimeoutSec))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "cannot sign url: %v", err)
	}

	return &blobpb.CreateBlobResponse{
		UploadUrl: u,
	}, nil
}

func secToDuration(sec int32) time.Duration {
	return time.Duration(sec) * time.Second
}
