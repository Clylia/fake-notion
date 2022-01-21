package page

import (
	"context"
	"fmt"
	pagepb "notion/page/api/gen/v1"
	"notion/page/page/dao"
	"notion/shared/auth"
	"notion/shared/id"
	"notion/shared/mongo/objid"

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

// GetPage gets a page.
func (s *Service) GetPage(c context.Context, req *pagepb.GetPageRequest) (*pagepb.PageEmtity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	page, err := s.Mongo.GetPage(c, id.PageID(req.Id), aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return page, nil
}

func (s *Service) GetPages(c context.Context, req *pagepb.GetPagesRequest) (*pagepb.GetPagesResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	pages, err := s.Mongo.GetPages(c, aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "cannot find any page")
	}
	return &pagepb.GetPagesResponse{
		Pages: pages,
	}, nil
}

// CreatePage creates a page.
func (s *Service) CreatePage(c context.Context, req *pagepb.CreatePageRequest) (*pagepb.PageEmtity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	blocks := convertBlockEmtity(req.Blocks)
	p, err := s.Mongo.CreatePage(c, id.AccountID(aid), blocks)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("account[%v] cannot create page, error: %+v", aid, err))
		return nil, status.Error(codes.Internal, "cannot create page")
	}
	return p, nil
}

// UpdatePage updates a page.
func (s *Service) UpdatePage(c context.Context, req *pagepb.UpdatePageRequest) (*pagepb.PageEmtity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	blocks := convertBlockEmtity(req.Blocks)
	p, err := s.Mongo.UpdatePage(c, id.PageID(req.Id), id.AccountID(aid), blocks)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("account[%v] cannot update page[%v], error: %+v", aid, req.Id, err))
		return nil, status.Error(codes.Internal, "cannot update page")
	}
	return p, nil
}

// DeletePage deletes a page.
func (s *Service) DeletePage(c context.Context, req *pagepb.DeletePageRequest) (*pagepb.DeletePageResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	err = s.Mongo.DeletePage(c, id.PageID(req.Id), aid)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("account[%v] delete page[%v] error", aid, req.Id), zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &pagepb.DeletePageResponse{}, nil
}

func convertBlockEmtity(blocks []*pagepb.BlockEmtity) []*dao.BlockEmtity {
	var bms []*dao.BlockEmtity
	for _, b := range blocks {
		var block dao.BlockEmtity
		block.ID = objid.MustFromID(id.BlockID(b.Id))
		block.HTML = b.Html
		block.Tag = b.Tag
		block.ImageURL = b.ImageUrl
		bms = append(bms, &block)
	}

	return bms
}
