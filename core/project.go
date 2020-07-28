// Package core ...
//
// Author: go_developer@163.com<张德满>
//
// Description: 项目信息维护
//
// File: project.go
//
// Version: 1.0.0
//
// Date: 2020/07/25 20:33:06
package core

import (
	"github.com/gin-gonic/gin"
	"github.com/go-developer/ginx-dao/define"
	"github.com/go-developer/go-util/util"
)

// NewProjectInstance 实例化项目管理
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 20:36:43
func NewProjectInstance() *Project {
	return &Project{}
}

// ExtraProjectInfo 项目扩展信息
//
// Author : go_developer@163.com<张德满>
type ExtraProjectInfo struct {
	Name         string
	Description  string
	CreateUserID uint64
}

// Project 项目管理
//
// Author : go_developer@163.com<张德满>
type Project struct{}

// Create 创建新项目
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 20:45:42
func (p *Project) Create(ctx *gin.Context, flag string, domain string, port uint, extra *ExtraProjectInfo) (*define.Project, error) {
	var (
		projectData *define.Project
		err         error
		//masterDBClient *godb.DBClient
		//slaveClient    *godb.DBClient
	)
	//masterDBClient = godb.DB.GetDBClient(ctx, false)
	projectData = p.buildProjectData(flag, domain, port, extra)
	//dao.NewProjectDao().Create()
	return projectData, err
}

// buildProjectData 构建项目数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 21:00:41
func (p *Project) buildProjectData(flag string, domain string, port uint, extra *ExtraProjectInfo) *define.Project {
	if nil == extra {
		extra = &ExtraProjectInfo{}
	}
	return &define.Project{
		ID:           util.ProjectUtil.GenerateID(),
		Flag:         flag,
		Domain:       domain,
		Port:         port,
		Name:         extra.Name,
		Description:  extra.Description,
		CreateUserID: extra.CreateUserID,
		ModifyUserID: extra.CreateUserID,
		CreateTime:   util.TimeUtil.GetCurrentFormatTime(),
		ModifyTime:   util.TimeUtil.GetCurrentFormatTime(),
	}
}
