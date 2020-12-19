// Package core...
//
// Description : core...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020-12-19 9:28 下午
package core

import (
	"github.com/gin-gonic/gin"
	"github.com/go-developer/ginx-dao/dao"
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

var (
	// Method 请求方法管理实例
	Method *method
)

func init()  {
	Method = &method{}
}

// methodList 请求方法列表
//
// Author : go_developer@163.com<张德满>
type methodList struct {
	List            []define.Method `json:"list"`
	Total           int64           `json:"total"`
	CurrentPage     int             `json:"current_page"`
	CurrentPageSize int64           `json:"current_page_size"`
}

type method struct {

}

// GetMethodByPage 分页获取method列表
//
// Author : go_developer@163.com<张德满>
func (m *method) GetMethodByPage(ctx *gin.Context, page int, size int64) (methodList, error) {
	var (
		err          error
		dbMethodList []define.Method
		dbClient     *godb.DBClient
		total        int64
	)
	dbClient = godb.DB.GetDBClient(ctx, true)
	if dbMethodList, err = dao.Method.GetMethodList(dbClient, dao.SetSearchOption{Func: dao.SetSearchOptionPage, Data: page}, dao.SetSearchOption{Func: dao.SetSearchOptionSize, Data: size}); nil != err {
		return methodList{List: make([]define.Method, 0), CurrentPage: page, CurrentPageSize: size}, err
	}
	if total, err = dao.Scheme.GetSchemeCount(dbClient); nil != err {
		return methodList{List: make([]define.Method, 0), CurrentPage: page, CurrentPageSize: size}, err
	}
	return methodList{
		List:            dbMethodList,
		Total:           total,
		CurrentPage:     page,
		CurrentPageSize: size,
	}, nil
}

