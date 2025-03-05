package service

import (
	"context"
	"service/internal/repo/clients"
	"service/internal/repo/db"
	"service/internal/service/entity"

	"github.com/aledenza/serco"
	"github.com/aledenza/serco/client"
	"github.com/aledenza/serco/database"
	"github.com/aledenza/serco/logger"
)

type UserServiceI interface {
	GetUser(ctx context.Context, userId int) (*entity.User, bool)
}

type UserService struct {
	client clients.UserClientI
	db     db.UserDatabaseI
}

func NewUserService(clientConfig client.ClientConfig, conn *database.Connection) *UserService {
	return &UserService{client: clients.NewUserClient(clientConfig), db: db.NewUserDatabase(conn)}
}

func (us *UserService) GetUser(ctx context.Context, userId int) (*entity.User, bool) {
	dbUser, success := us.db.GetUser(ctx, userId)
	if !success {
		return nil, false
	}
	clientUser, success := us.client.GetUser(ctx, userId)
	if !success {
		return nil, false
	}
	user, err := serco.Dump[entity.User](dbUser)
	if err != nil {
		logger.Error(ctx, "error scanning entity", err)
		return nil, false
	}
	if user.UserId != clientUser.UserId {
		logger.Error(ctx, "different entities", nil)
		return nil, false
	}
	return &user, true
}
