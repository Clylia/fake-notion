package account

import (
	"context"
	"notion/account/account/dao"
	accountpb "notion/account/api/gen/v1"
	"notion/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Encryptor defines enctyptor interface.
type Encryptor interface {
	Encrypt(password string) (hash string, err error)
}

// Service definds a car service.
type Service struct {
	accountpb.UnimplementedAccountServiceServer

	Encryptor Encryptor
	Mongo     *dao.Mongo
	Logger    *zap.Logger
}

// CreateAccount creates a account.
func (s *Service) CreateAccount(c context.Context, req *accountpb.CreateAccountRequest) (*accountpb.AccountEntity, error) {
	passowrd, err := s.Encryptor.Encrypt(req.Password)
	if err != nil {
		s.Logger.Error("cannot encrypt password", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	ar := &dao.AccountRecord{
		Account: &accountpb.Account{
			Username: req.Username,
			Email:    req.Email,
			Password: passowrd,
		},
	}
	ar, err = s.Mongo.CreateAccount(c, ar)
	if err != nil {
		s.Logger.Error("cannot create account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &accountpb.AccountEntity{
		Id:       ar.ID.Hex(),
		Username: ar.Account.Username,
		Email:    ar.Account.Email,
	}, nil
}

// GetAccount gets a account.
func (s *Service) GetAccount(c context.Context, req *accountpb.GetAccountRequest) (*accountpb.AccountEntity, error) {
	ar, err := s.Mongo.GetAccount(c, id.AccountID(req.Id))
	if err != nil {
		s.Logger.Error("cannot get account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &accountpb.AccountEntity{
		Id:       ar.ID.Hex(),
		Username: ar.Account.Username,
		Email:    ar.Account.Email,
	}, nil
}

// UpdateAccount updates a account.
func (s *Service) UpdateAccount(c context.Context, req *accountpb.UpdateAccountRequest) (*accountpb.UpdateAccountResponse, error) {
	update := &dao.AccountUpdate{
		Username: req.Username,
		Email:    req.Email,
	}
	_, err := s.Mongo.UpdateAccount(c, id.AccountID(req.Id), update)
	if err != nil {
		s.Logger.Error("cannot update account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &accountpb.UpdateAccountResponse{}, nil
}

// ChangePassword changes a account password.
func (s *Service) ChangePassword(c context.Context, req *accountpb.ChangePasswordRequest) (*accountpb.ChangePasswordResponse, error) {
	passowrd, err := s.Encryptor.Encrypt(req.Password)
	if err != nil {
		s.Logger.Error("cannot encrypt password", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	err = s.Mongo.ChangePassword(c, id.AccountID(req.Id), passowrd)
	if err != nil {
		s.Logger.Error("cannot change password", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &accountpb.ChangePasswordResponse{}, nil
}
