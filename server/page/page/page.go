package page

import (
	pagepb "notion/page/api/gen/v1"
	"notion/page/page/dao"

	"go.uber.org/zap"
)

// Service defines a page service.
type Service struct {
	pagepb.UnimplementedPageServiceServer

	Mongo  *dao.Mongo
	Logger *zap.Logger
}
