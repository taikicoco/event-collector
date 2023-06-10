package request

type GetWebLogsRequest struct {
}

type GetWebLogRequest struct {
	ID uint `param:"id"`
}

type CreateWebLogRequest struct {
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}

type UpdateWebLogRequest struct {
	ID      uint   `param:"id"`
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}
