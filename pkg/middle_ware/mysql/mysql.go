package mysql

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/grin-ch/grin-auth/pkg/model"
)

var Client *model.Client

func MysqlInit(dsn string) error {
	var err error
	Client, err = model.Open("mysql", dsn)
	if err != nil {
		return err
	}
	Client.Schema.Create(context.Background())
	return err
}
