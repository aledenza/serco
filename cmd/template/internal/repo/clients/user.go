package clients

import (
	"context"
	clientSchema "service/internal/repo/clients/schema"
	"strconv"

	"github.com/aledenza/serco/client"
	"github.com/aledenza/serco/logger"
)

type UserClientI interface {
	GetUser(ctx context.Context, userId int) (*clientSchema.User, bool)
}

type UserClient struct {
	client client.Client
}

func NewUserClient(config client.ClientConfig) *UserClient {
	return &UserClient{client: client.NewClient(config.URL, "UserClient", config)}
}

func (uc *UserClient) GetUser(ctx context.Context, userId int) (*clientSchema.User, bool) {
	response, err := uc.client.Get("GetUser")(ctx, "/api/v1/user/"+strconv.Itoa(userId))
	if err != nil {
		logger.Error(ctx, "error getting user", err)
		return nil, false
	}
	var user clientSchema.User
	if err = response.Scan(&user); err != nil {
		logger.Error(ctx, "error scanning response", err)
		return nil, false
	}
	return &user, true
}
