package account

import (
	"context"

	"github.com/grin-ch/grin-auth/pkg/model"
	"github.com/grin-ch/grin-auth/pkg/model/user"
	"github.com/grin-ch/grin-auth/pkg/service/internal"
)

type dbServer struct {
	client *model.Client
}

func newDbServer(client *model.Client) *dbServer {
	return &dbServer{
		client: client,
	}
}

// CreateUser 创建用户
func (s *dbServer) CreateUser(ctx context.Context,
	nickname, phone, email, passwd string) (*model.User, error) {
	return internal.Transaction(ctx, s.client, func(tx *model.Tx) (*model.User, error) {
		user, err := tx.User.Create().
			SetEmail(email).
			SetPhone(phone).
			SetPassword(passwd).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		_, err = tx.UserData.Create().
			SetNickname(nickname).
			SetUser(user).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return user, err
	})
}

// FindUserByAccount 查询用户
func (s *dbServer) FindUserByAccount(ctx context.Context, account string) (*model.User, error) {
	return s.client.User.Query().Where(user.Or(user.Email(account), user.Phone(account))).WithUserData().Only(ctx)
}
