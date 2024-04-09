package models

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page,optional" `    // 页码
	PageSize int    `json:"pageSize,optional"` // 每页大小
	Keyword  string `json:"keyword,optional" ` //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
