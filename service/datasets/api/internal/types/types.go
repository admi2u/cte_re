// Code generated by goctl. DO NOT EDIT.
package types

type Dataset struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	DatasetSize int64  `json:"dataset_size"`
	DatasetType string `json:"dataset_type"`
	CreateAt    string `json:"create_at"`
	UpdateAt    string `json:"update_at"`
}

type DatasetCreateRequest struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	DatasetSize int64  `json:"dataset_size"`
	DatasetType string `json:"dataset_type"`
}

type DatasetCreateResponse struct {
	Id int64 `json:"id"`
}

type DatasetListRequest struct {
	Page     int64 `form:"page"`
	PageSize int64 `form:"page_size"`
}

type DatasetListResponse struct {
	Total    int64      `json:"total"`
	Datasets []*Dataset `json:"datasets"`
}