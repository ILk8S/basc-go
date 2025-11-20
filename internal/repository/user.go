package repository

import (
	"context"
	"errors"

	"github.com/ILk8S/basc-go/internal/domain"
	"github.com/ILk8S/basc-go/internal/repository/cache"
	"github.com/ILk8S/basc-go/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO, c *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: c,
	}
}

func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Passowrd: u.Password,
	})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := repo.cache.Get(ctx, id)
	if err == nil {
		return u, err
	}
	//如果redis没查到，去查数据库
	ue, err := repo.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	u = repo.toDomain(ue)
	// 忽略了这个错误，因为不要因为放不进redis而阻塞住
	_ = repo.cache.Set(ctx, u)
	return u, nil
}

func (repo *UserRepository) FindByIdV1(ctx context.Context, id int64) (domain.User, error) {
	u, err := repo.cache.Get(ctx, id)
	switch {
	case err == nil:
		return u, err
	case errors.Is(err, cache.ErrKeyNotExist):
		ue, err := repo.dao.FindById(ctx, id)
		if err != nil {
			return domain.User{}, err
		}
		u = repo.toDomain(ue)
		_ = repo.cache.Set(ctx, u)
		return u, nil
	default:
		return domain.User{}, err
	}
}

func (repo *UserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Passowrd,
	}
}
