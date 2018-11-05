package entity

type Answer struct {
	Id int64 `json:"id"`
	Content string `json:"content"`
}

type PagingAnswer struct {
	Data []Answer `json:"data"`
	Paging Paging `json:"paging"`
}
