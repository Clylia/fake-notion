package auth

import (
	"context"
	"fmt"
	"notion/auth/auth/dao"
	"notion/shared/id"
	"time"

	authpb "notion/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements auth service.
type Service struct {
	authpb.UnimplementedAuthServiceServer

	Decryptor          Decryptor
	TokenGenerator     TokenGenerator
	TokenVerifier      TokenVerifier
	AccessTokenExprie  time.Duration
	RefreshTokenExprie time.Duration
	Monogo             *dao.Mongo
	Logger             *zap.Logger
}

// Decryptor defines dnctyptor interface.
type Decryptor interface {
	Compare(password, hash string) (ok bool, err error)
}

// TokenGenerator generates a token for the specified account.
type TokenGenerator interface {
	GenAccessToken(accountID string, exprie time.Duration) (string, error)
	GenRefreshToken(accountID string, exprie time.Duration) (string, error)
}

// TokenVerifier verifies a token for the specified account.
type TokenVerifier interface {
	VerifyRefreshToken(token string) (aid string, err error)
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
		s.Logger.Error("cannot compare login password", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	if !ok {
		return nil, status.Error(codes.Aborted, fmt.Sprintf("login email[%v] password invalid", req.Email))
	}

	accTkn, err := s.TokenGenerator.GenAccessToken(ar.ID.String(), s.AccessTokenExprie)
	if err != nil {
		s.Logger.Error("cannot generate access token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	refTkn, err := s.TokenGenerator.GenRefreshToken("", s.RefreshTokenExprie)
	if err != nil {
		s.Logger.Error("cannot generate refresh token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken:  accTkn,
		RefreshToken: refTkn,
		ExpiresIn:    int32(time.Now().Add(s.AccessTokenExprie).Unix()),
	}, nil
}

// Refresh refreshs login.
func (s *Service) Refresh(c context.Context, req *authpb.RefreshLoginRequest) (*authpb.RefreshLoginResponse, error) {
	aid, err := s.TokenVerifier.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	_, err = s.Monogo.Exists(c, id.AccountID(aid))
	if err != nil {
		s.Logger.Error(fmt.Sprintf("account id[%v] does no exists", aid), zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "")
	}
	accTkn, err := s.TokenGenerator.GenAccessToken(aid, s.AccessTokenExprie)
	if err != nil {
		s.Logger.Error("cannot generate access token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	refTkn, err := s.TokenGenerator.GenRefreshToken(aid, s.RefreshTokenExprie)
	if err != nil {
		s.Logger.Error("cannot generate refresh token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.RefreshLoginResponse{
		AccessToken:  accTkn,
		RefreshToken: refTkn,
		ExpiresIn:    int32(s.AccessTokenExprie.Seconds()),
	}, nil
}
