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
	"errors"

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

// schemeList 分页查询,输出的数据结构
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/08 01:16:13
type schemeList struct {
	List            []define.Scheme `json:"list"`
	Total           int64           `json:"total"`
	CurrentPage     int             `json:"current_page"`
	CurrentPageSize int64           `json:"current_page_size"`
}

func init() {
	Scheme = &scheme{}
}

type scheme struct{}

// CreateScheme 创建一个scheme
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/30 23:32:55
func (s *scheme) CreateScheme(ctx *gin.Context, scheme string) (*define.Scheme, error) {
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

// SoftDelete 软删除scheme
//
// Author : go_developer@163.com<张德满>
func (s *scheme) SoftDelete(ctx *gin.Context, schemeID uint64, currentStatus uint) error {
	dbClient := godb.DB.GetDBClient(ctx, false)
	if affectRows, err := dao.Scheme.ChangeStatus(dbClient, schemeID, currentStatus, define.SchemeStatusForbbiden); nil != err || affectRows < 1 {
		return errors.New("更新失败, scheme不存在或状态异常")
	}
	return nil
}

// GetAllScheme 获取全部scheme列表
//
// Author : go_developer@163.com<张德满>
func (s *scheme) GetAllScheme(ctx *gin.Context) ([]*define.Scheme, error) {
	dbClient := godb.DB.GetDBClient(ctx, true)
	return dao.Scheme.GetAllScheme(dbClient)
}

// UpdateScheme 更新scheme
//
// Author : go_developer@163.com<张德满>
func (s *scheme) UpdateScheme(ctx *gin.Context, schemeID uint64, schemeName string) error {
	dbClient := godb.DB.GetDBClient(ctx, false)
	if affectRows, err := dao.Scheme.Update(dbClient, schemeID, schemeName); affectRows < 1 || nil != err {
		return errors.New("更新scheme信息失败")
	}
	return nil
}

// GetSchemeByPage 分页获取scheme列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/08 01:17:33
func (s *scheme) GetSchemeByPage(ctx *gin.Context, page int, size int64) (schemeList, error) {
	var (
		err          error
		dbSchemeList []define.Scheme
		dbClient     *godb.DBClient
		total        int64
	)
	dbClient = godb.DB.GetDBClient(ctx, false)
	if dbSchemeList, err = dao.Scheme.GetSchemeList(dbClient, dao.SetSearchOption{Func: dao.SetSearchOptionPage, Data: page}, dao.SetSearchOption{Func: dao.SetSearchOptionSize, Data: size}); nil != err {
		return schemeList{List: make([]define.Scheme, 0), CurrentPage: page, CurrentPageSize: size}, err
	}
	if total, err = dao.Scheme.GetSchemeCount(dbClient); nil != err {
		return schemeList{List: make([]define.Scheme, 0), CurrentPage: page, CurrentPageSize: size}, err
	}
	return schemeList{
		List:            dbSchemeList,
		Total:           total,
		CurrentPage:     page,
		CurrentPageSize: size,
	}, nil
}
