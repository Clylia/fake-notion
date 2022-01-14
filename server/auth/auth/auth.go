package auth

import (
	"context"
	"fmt"
	"notion/auth/auth/dao"
	"time"

	authpb "notion/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements auth service.
type Service struct {
	authpb.AuthServiceServer
	Decryptor      Decryptor
	TokenGenerator TokenGenerator
	TokenExprie    time.Duration
	Monogo         *dao.Mongo
	Logger         *zap.Logger
}

// Decryptor defines dnctyptor interface.
type Decryptor interface {
	Compare(password, hash string) (ok bool, err error)
}

// TokenGenerator generates a token for the specified account.
type TokenGenerator interface {
	GenerateToken(accountID string, exprie time.Duration) (string, error)
}

// Login logs a user in.
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	ar, err := s.Monogo.ResolveAccount(c, req.Email)
	if err != nil {
		s.Logger.Error("cannot resolve account", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	ok, err := s.Decryptor.Compare(req.Password, ar.Account.Password)
	if err != nil {
		s.Logger.Error("cannot compare account password", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	if !ok {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("login email[%v] password invalid", req.Email))
	}

	tkn, err := s.TokenGenerator.GenerateToken(ar.ID.String(), s.TokenExprie)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresIn:   int32(s.TokenExprie.Seconds()),
	}, nil
}
