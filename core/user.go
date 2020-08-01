// Package core ...
//
// Author: go_developer@163.com<张德满>
//
// Description: 用户管理逻辑
//
// File: user.go
//
// Version: 1.0.0
//
// Date: 2020/08/01 20:32:58
package core

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/ginx-dao/dao"
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

var user *User

// NewCoreUser 获取实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 22:29:53
func NewCoreUser() *User {
	if nil == user {
		user = &User{}
	}
	return user
}

// User 用户管理
//
// Author : go_developer@163.com<张德满>
type User struct{}

// GetUserDatailByPhone 获取用户详情
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 20:37:36
func (u *User) GetUserDatailByPhone(ctx *gin.Context, phone string) (*define.User, error) {
	var (
		dbClient     *godb.DBClient
		err          error
		userDao      *dao.UserDao
		userInfo     *define.User
		userNotExist bool
	)
	userDao = dao.NewUserDao()
	dbClient = godb.DB.GetDBClient(ctx, true)
	userInfo, userNotExist, err = userDao.GetDetailByPhone(dbClient, phone)
	if userNotExist {
		return nil, errors.New("用户不存在")
	}
	return userInfo, err
}

// GetUserDatailByID 获取用户详情
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 20:37:36
func (u *User) GetUserDatailByID(ctx *gin.Context, userID uint64) (*define.User, error) {
	var (
		dbClient     *godb.DBClient
		err          error
		userDao      *dao.UserDao
		userInfo     *define.User
		userNotExist bool
	)
	userDao = dao.NewUserDao()
	dbClient = godb.DB.GetDBClient(ctx, true)
	userInfo, userNotExist, err = userDao.GetDetailByID(dbClient, userID)
	if userNotExist {
		return nil, errors.New("用户不存在")
	}
	return userInfo, err
}
