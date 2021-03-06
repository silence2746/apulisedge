// Copyright 2020 Apulis Technology Inc. All rights reserved.

package imageservice

import (
	apulisdb "github.com/apulis/ApulisEdge/cloud/pkg/database"
	imagemodule "github.com/apulis/ApulisEdge/cloud/pkg/domain/image"
	imageentity "github.com/apulis/ApulisEdge/cloud/pkg/domain/image/entity"
	proto "github.com/apulis/ApulisEdge/cloud/pkg/protocol"
	"strings"
)

// list container image version
func ListContainerImageVersion(userInfo proto.ApulisHeader, req *imagemodule.ListContainerImageVersionReq) (*[]imageentity.UserContainerImageVersionInfo, int, error) {
	var imageVerInfos []imageentity.UserContainerImageVersionInfo

	total := 0
	offset := req.PageSize * (req.PageNum - 1)
	limit := req.PageSize

	res := apulisdb.Db.Offset(offset).Limit(limit).
		Where("ClusterId = ? and GroupId = ? and UserId = ? and ImageName = ? and OrgName = ?",
			userInfo.ClusterId, userInfo.GroupId, userInfo.UserId, req.ImageName, req.OrgName).
		Find(&imageVerInfos)

	if res.Error != nil {
		return &imageVerInfos, total, res.Error
	}

	total = int(res.RowsAffected)
	return &imageVerInfos, total, nil
}

// describe container image version
func DescribeContainerImageVersion(userInfo proto.ApulisHeader, req *imagemodule.DescribeContainerImageVersionReq) (*imageentity.UserContainerImageVersionInfo, error) {
	var imageVerInfo imageentity.UserContainerImageVersionInfo

	res := apulisdb.Db.
		Where("ClusterId = ? and GroupId = ? and UserId = ? and ImageName = ? and OrgName = ? and ImageVersion = ?",
			userInfo.ClusterId, userInfo.GroupId, userInfo.UserId, req.ImageName, req.OrgName, req.ImageVersion).
		First(&imageVerInfo)

	if res.Error != nil {
		return &imageVerInfo, res.Error
	}

	return &imageVerInfo, nil
}

// delete container image version
func DeleteContainterImageVersion(userInfo proto.ApulisHeader, req *imagemodule.DeleteContainerImageVersionReq) error {
	var imageVerInfo imageentity.UserContainerImageVersionInfo

	// get image and delete
	res := apulisdb.Db.
		Where("ClusterId = ? and GroupId = ? and UserId = ? and ImageName = ? and OrgName = ? and ImageVersion = ?",
			userInfo.ClusterId, userInfo.GroupId, userInfo.UserId, req.ImageName, req.OrgName, req.ImageVersion).
		First(&imageVerInfo)
	if res.Error != nil {
		return res.Error
	}

	err := imageentity.DeleteContainerImageVersion(&imageVerInfo)
	if err != nil {
		return err
	}

	return nil
}

func DoIHaveTheImageVersion(userInfo proto.ApulisHeader, orgName string, imgName string, imgVersion string) (string, bool) {
	var verInfo imageentity.UserContainerImageVersionInfo

	res := apulisdb.Db.Model(&imageentity.UserContainerImageVersionInfo{}).
		Where("ClusterId = ? and GroupId = ? and UserId = ? and OrgName = ? and ImageName = ? and ImageVersion = ?",
			userInfo.ClusterId, userInfo.GroupId, userInfo.UserId, orgName, imgName, imgVersion).
		First(&verInfo)
	if res.Error == nil {
		return getImgPathFromDownloadCommand(verInfo.DownloadCommand), true
	} else {
		return "", false
	}
}

func getImgPathFromDownloadCommand(downloadCommand string) string {
	return strings.TrimPrefix(strings.Split(downloadCommand, imagemodule.DockerPullPrefix)[1], imagemodule.BlankString)
}
