package auth

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-developer/ginx-dao/define"
	"github.com/go-developer/ginx-manager/core"
	"github.com/go-developer/go-util/util"
)

var login *Login

// NewLogin 获取登录实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 21:44:34
func NewLogin() *Login {
	if nil == login {
		login = &Login{}
	}
	return login
}

// Login ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 21:45:22
type Login struct{}

// Login 登录
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 22:28:23
func (l *Login) Login(ctx *gin.Context, account string, password string) (string, error) {
	var (
		userInfo *define.User
		err      error
		token    string
	)
	if userInfo, err = core.NewCoreUser().GetUserDatailByPhone(ctx, account); nil != err {
		return "", err
	}
	token = fmt.Sprintf("%d", userInfo.ID)
	return token, nil
}

// Verify 验证token
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 22:37:36
func (l *Login) Verify(ctx *gin.Context, token string) (*define.User, error) {
	var (
		userID uint64
		err    error
	)
	if err = util.ConvertAssign(&userID, token); nil != err {
		return nil, errors.New("非法的token")
	}

	return core.NewCoreUser().GetUserDatailByID(ctx, userID)
}
