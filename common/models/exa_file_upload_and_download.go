package models

import (
	"gorm.io/gorm"
)

type ExaFileUploadAndDownload struct {
	gorm.Model
	FileName         string `json:"filename" gorm:"comment:文件名"`           // 文件名
	FileUrl          string `json:"fileurl" gorm:"comment:文件地址"`           // 文件地址
	FileTag          string `json:"filetag" gorm:"comment:文件标签"`           // 文件标签
	FileKey          string `json:"filekey" gorm:"comment:编号"`             // 编号
	FileUploaderName string `json:"fileuploadername" gorm:"comment:上传者账号"` // 上传者账号
}

// 上传小文件时需携带参数
type UploadFileReq struct {
	// File []byte `form:"file"`
	Path string `form:"path"`
}

type UploadBigFileReq struct {
	//File     multipart.File
	Flag     string `form:"flag,optional"`
	Hash     string `form:"hash,optional"`
	FileName string `form:"filename,optional"`
	FilePath string `form:"filepath,optional"`
	Index    int64  `form:"index,optional"`
	Total    int64  `form:"total,optional"`
	Size     int64  `form:"size,optional"`
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
