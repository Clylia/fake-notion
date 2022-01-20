package page

import (
	"context"
	pagepb "notion/page/api/gen/v1"
	"notion/page/page/dao"
	"notion/shared/auth"
	"notion/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	page := dao.ConvertPageRecord(req.Page)

	p, err := s.Mongo.CreatePage(c, id.AccountID(aid), page)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create page: %+v", err)
	}
	return p, nil
}

func (s *Service) UpdatePage(c context.Context, req *pagepb.UpdatePageRequest) (*pagepb.PageEmtity, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) DeletePage(c context.Context, req *pagepb.DeletePageRequest) (*pagepb.DeletePageResponse, error) {
	panic("not implemented") // TODO: Implement
}
