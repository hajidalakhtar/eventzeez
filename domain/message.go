package domain

type WebResponse struct {
	Code     int               `json:"code"`
	Status   string            `json:"status"`
	Data     interface{}       `json:"data"`
	Paginate PaginatedResponse `json:"paginate"`
}
