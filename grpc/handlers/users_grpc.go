package handlers

import (
	"context"

	"github.com/sunnysingha911/user-service/gen/user-service/userpb"
	"github.com/sunnysingha911/user-service/models"
	"github.com/sunnysingha911/user-service/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toProtoUser(user *models.User) *userpb.User {
	return &userpb.User{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt.String(),
		CreatedAt: user.CreatedAt.String(),
	}
}

func toProtoUsers(users []*models.User) []*userpb.User {
	var pbUsers []*userpb.User
	for _, u := range users {
		pbUsers = append(pbUsers, toProtoUser(u))
	}
	return pbUsers
}

func (s *AuthServiceServer) GetAllUsers(ctx context.Context, req *userpb.GetAllUserRequest) (*userpb.UsersListResponse, error) {
	res, err := services.GetUserList(int(req.Page), int(req.Limit))
	if err == services.ErrInvalidCredentials {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &userpb.UsersListResponse{
		Users: toProtoUsers(res.Users),
		Meta: &userpb.Meta{
			Limit: int32(res.Limit),
			Page:  int32(res.Page),
			Total: int32(res.Total),
		},
	}, nil
}

func (s *AuthServiceServer) GetUserById(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	res, err := services.GetUserById(req.Id)
	if err == services.ErrInvalidCredentials {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &userpb.GetUserResponse{
		User: toProtoUser(res.User),
	}, nil
}
