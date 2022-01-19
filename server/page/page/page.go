package page

import (
	"context"
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

func (s *Service) GetPage(c context.Context, req *pagepb.GetPageRequest) (*pagepb.PageEmtity, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) GetPages(c context.Context, req *pagepb.GetPagesRequest) (*pagepb.GetPagesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) CreatePage(c context.Context, req *pagepb.CreatePageRequest) (*pagepb.PageEmtity, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) UpdatePage(c context.Context, req *pagepb.UpdatePageRequest) (*pagepb.PageEmtity, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) DeletePage(c context.Context, req *pagepb.DeletePageRequest) (*pagepb.DeletePageResponse, error) {
	panic("not implemented") // TODO: Implement
}
