// Code generated by goctl. DO NOT EDIT.
package types

type GetAllImageResp struct {
	ImageList interface{} `json:"imagelist"`
}

type GetMyImageResp struct {
	ImageList interface{} `json:"imagelist"`
}

type GetUserImageReq struct {
	UserId int64 `json:"userid"`
}

type GetUserImageResp struct {
	ImageList interface{} `json:"imagelist"`
}