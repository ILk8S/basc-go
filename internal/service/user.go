package service

import (
	"context"

	"github.com/ILk8S/basc-go/internal/domain"
)


type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

func (svc *UserService) Signup(ctx context.Context, u *domain.User) error {
	return nil
}

