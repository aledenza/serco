package db

import (
	"context"

	"github.com/aledenza/serco/database"
	"github.com/aledenza/serco/logger"
)

type UserDatabaseI interface {
	GetUser(ctx context.Context, userId int) (map[string]any, bool)
}

type UserDatabase struct {
	db database.Database
}

func NewUserDatabase(conn *database.Connection) *UserDatabase {
	db := database.NewDatabase("User", conn)
	return &UserDatabase{db: db}
}

func (ud *UserDatabase) GetUser(ctx context.Context, userId int) (map[string]any, bool) {
	query := `
		SELECT user_id, first_name, second_name
		FROM user_table
		WHERE user_id = @user_id
	`
	params := map[string]any{
		"user_id": userId,
	}
	result, err := ud.db.FetchOne("GetUser")(ctx, query, params)
	if err != nil {
		logger.Error(ctx, "Error getting user", err)
		return nil, false
	}
	return result, true
}
