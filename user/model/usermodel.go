package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel

		// username 无缓存查找数据
		FindOneByUsername(ctx context.Context, username string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

// username 无缓存查找数据
func (m *customUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	querysql := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)

	err := m.QueryRowNoCache(&user, querysql, username)
	if err != nil {
		logx.Errorf("username.findOne error, username=%s, err=%s", username, err.Error())
		if err == sqlx.ErrNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
