// Package core ...
//
// Author: go_developer@163.com<张德满>
//
// Description: 管理scheme & method
//
// File: scheme.go
//
// Version: 1.0.0
//
// Date: 2020/08/30 23:25:38
package core

import (
	"github.com/gin-gonic/gin"
	"github.com/go-developer/ginx-dao/dao"
	"github.com/go-developer/ginx-dao/define"
	"github.com/go-developer/go-util/util"
	godb "github.com/go-developer/gorm-mysql"
)

var (
	// Scheme 操作实例
	Scheme *scheme
)

func init() {
	Scheme = &scheme{}
}

type scheme struct{}

// CreateScheme 创建一个scheme
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/30 23:32:55
func (s *scheme) CreateScheme(ctx *gin.Context, scheme string) (uint64, error) {
	return 0, nil
}

// CreateMethod 创建请求方法
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/30 23:34:26
func (s *scheme) CreateMethod(ctx *gin.Context, scheme string) (*define.Scheme, error) {
	schemeData := &define.Scheme{
		Scheme:       scheme,
		Status:       define.APIParamStatusUsing,
		CreateUserID: 0,
		ModifyUserID: 0,
		CreateTime:   util.TimeUtil.GetCurrentFormatTime(),
		ModifyTime:   util.TimeUtil.GetCurrentFormatTime(),
	}
	dbClient := godb.DB.GetDBClient(ctx, false)
	if err := dao.Scheme.CreateScheme(dbClient, schemeData); nil != err {
		return nil, err
	}
	return schemeData, nil
}

// GetAllScheme 获取全部scheme列表
//
// Author : go_developer@163.com<张德满>
func (s *scheme) GetAllScheme(ctx *gin.Context) ([]*define.Scheme, error) {
	dbClient := godb.DB.GetDBClient(ctx, true)
	return dao.Scheme.GetAllScheme(dbClient)
}
