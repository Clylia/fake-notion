package account

import (
	"context"
	"fmt"
	"notion/account/account/dao"
	accountpb "notion/account/api/gen/v1"
	"notion/shared/auth"
	"notion/shared/errs"
	"notion/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Encryptor defines enctyptor interface.
type Encryptor interface {
	Encrypt(password string) (hash string, err error)
}

// Service defines a car service.
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
		if errs.IsDuplicateKeyErr(err) {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("username[%v] or email[%v] has taken.", req.Username, req.Email))
		}
		s.Logger.Error("cannot create account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	ar.Account.Password = ""
	return &accountpb.AccountEntity{
		Id:      ar.ID.Hex(),
		Account: ar.Account,
	}, nil
}

// GetAccount gets a account.
func (s *Service) GetAccount(c context.Context, req *accountpb.GetAccountRequest) (*accountpb.AccountEntity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	if aid != id.AccountID(req.Id) {
		s.Logger.Error(fmt.Sprintf("login account[%v] cannot get another account[%v]", aid.String(), req.Id))
		return nil, status.Error(codes.Unauthenticated, "")
	}

	ar, err := s.Mongo.GetAccount(c, aid)
	if err != nil {
		s.Logger.Error("cannot get account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	ar.Account.Password = ""
	return &accountpb.AccountEntity{
		Id:      ar.ID.Hex(),
		Account: ar.Account,
	}, nil
}

// UpdateAccount updates a account.
func (s *Service) UpdateAccount(c context.Context, req *accountpb.UpdateAccountRequest) (*accountpb.UpdateAccountResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	if aid != id.AccountID(req.Id) {
		s.Logger.Error(fmt.Sprintf("login account[%v] cannot update another account[%v]", aid.String(), req.Id))
		return nil, status.Error(codes.Unauthenticated, "")
	}
	update := &dao.AccountUpdate{
		Username: req.Username,
		Email:    req.Email,
	}
	_, err = s.Mongo.UpdateAccount(c, aid, update)
	if err != nil {
		if errs.IsDuplicateKeyErr(err) {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("username[%v] or email[%v] has taken.", req.Username, req.Email))
		}
		s.Logger.Error("cannot update account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &accountpb.UpdateAccountResponse{}, nil
}

// ChangePassword changes a account password.
func (s *Service) ChangePassword(c context.Context, req *accountpb.ChangePasswordRequest) (*accountpb.ChangePasswordResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	if aid != id.AccountID(req.Id) {
		s.Logger.Error(fmt.Sprintf("login account[%v] cannot update another account[%v]", aid.String(), req.Id))
		return nil, status.Error(codes.Unauthenticated, "")
	}
	passowrd, err := s.Encryptor.Encrypt(req.Password)
	if err != nil {
		s.Logger.Error("cannot encrypt password", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = s.Mongo.ChangePassword(c, aid, passowrd)
	if err != nil {
		s.Logger.Error("cannot change password", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &accountpb.ChangePasswordResponse{}, nil
}
