package user

import (
	"context"
	"fmt"

	"github.com/22maksim/family_app/api/proto/user"
)

type userServiceServerImpl struct {
	user.UnimplementedUserServiceServer
}

func (s *userServiceServerImpl) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	fmt.Println("New user:", req.Name, req.Email)
	return &user.RegisterResponse{UserId: "12345"}, nil
}

func (s *userServiceServerImpl) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	// Логика логина
	return &user.LoginResponse{Token: "ugtwdv13532kkj233324"}, nil
}
