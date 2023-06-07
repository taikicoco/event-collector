package request

type UpdateWebLogDataRequest struct {
	ID         uint `json:"id"`
	WebLogID   uint `json:"web_log_id"`
	Access     uint `json:"access"`
	Conversion uint `json:"conversion"`
}
