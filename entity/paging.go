package entity

type Paging struct {
	IsEnd bool `json:"is_end"`
	IsStart bool `json:"is_start"`
	Total int `json:"total"`
}
