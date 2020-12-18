// Copyright 2020 Apulis Technology Inc. All rights reserved.

package image

import (
	imageentity "github.com/apulis/ApulisEdge/cloud/pkg/domain/image/entity"
	"mime/multipart"
	"time"
)

// TODO add param validate, like node.ListEdgeNodesReq

// List image
type ListContainerImageReq struct {
	PageNum  int `json:"pageNum" validate:"gte=1,lte=1000"`
	PageSize int `json:"pageSize" validate:"gte=1,lte=1000"`
}

type RspContainerImageInfo struct {
	ClusterId    int64     `json:"clusterId"`
	GroupId      int64     `json:"groupId"`
	UserId       int64     `json:"userId"`
	ImageName    string    `json:"imageName"`
	OrgName      string    `json:"orgName"`
	VersionCount int       `json:"versionCount"`
	UpdateAt     time.Time `json:"updateAt"`
}

type ListContainerImageRsp struct {
	Total  int                     `json:"total"`
	Images []RspContainerImageInfo `json:"images"`
}

// Upload image
type UploadContainerImageReq struct {
	File    *multipart.FileHeader `form:"file" binding:"required"`
	OrgName string                `form:"orgName" binding:"required"`
}

type UploadContainerImageRsp struct {
}

// Delete image
type DeleteContainerImageReq struct {
	ImageName string `json:"imageName"`
	OrgName   string `json:"orgName"`
}

type DeleteContainerImageRsp struct {
}

// List image version
type ListContainerImageVersionReq struct {
	ImageName string `json:"imageName"`
	OrgName   string `json:"orgName"`
	PageNum   int    `json:"pageNum" validate:"gte=1,lte=1000"`
	PageSize  int    `json:"pageSize" validate:"gte=1,lte=1000"`
}

type ListContainerImageVersionRsp struct {
	Total         int                                          `json:"total"`
	ImageVersions *[]imageentity.UserContainerImageVersionInfo `json:"imageVersions"`
}

// Delete image version
type DeleteContainerImageVersionReq struct {
	ImageName    string `json:"imageName"`
	OrgName      string `json:"orgName"`
	ImageVersion string `json:"imageVersion"`
}

type DeleteContainerImageVersionRsp struct {
}

// Create image org
type CreateContainerImageOrgReq struct {
	OrgName string `json:"orgName"`
}

type CreateContainerImageOrgRsp struct {
	Org *imageentity.ContainerImageOrg `json:"org"`
}

// List image org
type ListContainerImageOrgReq struct {
	PageNum  int `json:"pageNum" validate:"gte=1,lte=1000"`
	PageSize int `json:"pageSize" validate:"gte=1,lte=1000"`
}

type ListContainerImageOrgRsp struct {
	Total     int                              `json:"total"`
	ImageOrgs *[]imageentity.ContainerImageOrg `json:"imageOrgs"`
}

// Delete image org
type DeleteContainerImageOrgReq struct {
	OrgName string `json:"orgName"`
}

type DeleteContainerImageOrgRsp struct {
}
