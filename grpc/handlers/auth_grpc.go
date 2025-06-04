package handlers

import (
	"context"

	"github.com/sunnysingha911/user-service/gen/user-service/userpb"
	"github.com/sunnysingha911/user-service/models"
	"github.com/sunnysingha911/user-service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

func buildProtoUser(user *models.User) *userpb.User {
	return &userpb.User{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt.String(),
		CreatedAt: user.CreatedAt.String(),
	}
}

func (s *AuthServiceServer) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	res, err := services.Login(req.GetEmail(), req.GetPassword())
	if err == services.ErrInvalidCredentials {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &userpb.LoginResponse{
		Token: res.Token,
		User:  buildProtoUser(res.User),
	}, nil
}
