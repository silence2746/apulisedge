// Copyright 2020 Apulis Technology Inc. All rights reserved.

package httpserver

import (
	appmodule "github.com/apulis/ApulisEdge/cloud/pkg/domain/application"
	appentity "github.com/apulis/ApulisEdge/cloud/pkg/domain/application/entity"
	appservice "github.com/apulis/ApulisEdge/cloud/pkg/domain/application/service"
	proto "github.com/apulis/ApulisEdge/cloud/pkg/protocol"
	"github.com/gin-gonic/gin"
)

func ApplicationHandlerRoutes(r *gin.Engine) {
	group := r.Group("/apulisEdge/api/application")

	// add authentication
	group.Use(Auth())

	/// edge application
	group.POST("/listApplication", wrapper(ListEdgeApps))
	group.POST("/createApplication", wrapper(CreateEdgeApp))
	group.POST("/describeApplication", wrapper(DescribeEdgeApp))
	group.POST("/deleteApplication", wrapper(DeleteEdgeApp))

	// edge application version
	group.POST("/listApplicationVersion", wrapper(ListEdgeAppVersions))
	group.POST("/describeApplicationVersion", wrapper(DescribeEdgeAppVersion))
	group.POST("/deleteApplicationVersion", wrapper(DeleteEdgeAppVersion))
	group.POST("/publishApplicationVersion", wrapper(PublishEdgeAppVersion))
	group.POST("/offlineApplicationVersion", wrapper(OfflineEdgeAppVersion))

	// application deployment
	group.POST("/listApplicationDeploy", wrapper(ListEdgeAppDeploys))
	group.POST("/listNodeDeploy", wrapper(ListNodeDeploys))
	group.POST("/deployApplication", wrapper(DeployEdgeApp))
	group.POST("/undeployApplication", wrapper(UnDeployEdgeApp))
	group.POST("/listCanDeployNode", wrapper(ListNodeCanDeploy))
}

// list edge apps
func ListEdgeApps(c *gin.Context) error {
	var err error
	var req proto.Message
	var reqContent appmodule.ListEdgeApplicationReq
	var apps *[]appentity.ApplicationBasicInfo
	var total int

	userInfo, errRsp := PreHandler(c, &req, &reqContent)
	if errRsp != nil {
		return errRsp
	}

	// list node
	apps, total, err = appservice.ListEdgeApplications(*userInfo, &reqContent)
	if err != nil {
		return AppError(c, &req, APP_ERROR_CODE, err.Error())
	}

	data := appmodule.ListEdgeApplicationRsp{
		Total: total,
		Apps:  apps,
	}
	return SuccessResp(c, &req, data)
}

// create edge application
// this interface can both create basic app and app version
func CreateEdgeApp(c *gin.Context) error {
	var err error
	var req proto.Message
	var reqContent appmodule.CreateEdgeApplicationReq

	userInfo, errRsp := PreHandler(c, &req, &reqContent)
	if errRsp != nil {
		return errRsp
	}

	// create application
	appCreated, verCreated, err := appservice.CreateEdgeApplication(*userInfo, &reqContent)
	if err != nil {
		return AppError(c, &req, APP_ERROR_CODE, err.Error())
	}

	data := appmodule.CreateEdgeApplicationRsp{
		AppCreated:     appCreated,
		VersionCreated: verCreated,
	}

	return SuccessResp(c, &req, data)
}

// describe edge application
func DescribeEdgeApp(c *gin.Context) error {
	var err error
	var req proto.Message
	var reqContent appmodule.DescribeEdgeApplicationReq

	userInfo, errRsp := PreHandler(c, &req, &reqContent)
	if errRsp != nil {
		return errRsp
	}

	// describe application
	app, err := appservice.DescribeEdgeApp(*userInfo, &reqContent)
	if err != nil {
		return AppError(c, &req, APP_ERROR_CODE, err.Error())
	}

	data := appmodule.DescribeEdgeApplicationRsp{
		App: app,
	}

	return SuccessResp(c, &req, data)
}

// delete edge application
func DeleteEdgeApp(c *gin.Context) error {
	var err error
	var req proto.Message
	var reqContent appmodule.DeleteEdgeApplicationReq

	userInfo, errRsp := PreHandler(c, &req, &reqContent)
	if errRsp != nil {
		return errRsp
	}

	// delete application
	err = appservice.DeleteEdgeApplication(*userInfo, &reqContent)
	if err != nil {
		return AppError(c, &req, APP_ERROR_CODE, err.Error())
	}

	return SuccessResp(c, &req, "OK")
}
