package request

type GetWebLogsRequest struct {
}

type GetWebLogRequest struct {
	ID uint `json:"id"`
}

type CreateWebLogRequest struct {
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}

type UpdateWebLogRequest struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Version uint   `json:"version"`
	PageUrl string `json:"page_url"`
}
