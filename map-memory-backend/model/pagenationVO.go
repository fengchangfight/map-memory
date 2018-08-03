package model

type PaginationVO struct {
	TotalCount  int         `json:"totalCount"`
	PageSize    int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Data        interface{} `json:"data"`
}
