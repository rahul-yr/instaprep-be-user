package helpers

type ErrorResponse struct {
	Error  string `json:"error"`
	Status bool   `json:"status"`
}
