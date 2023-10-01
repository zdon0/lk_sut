package dto

type SimpleOkResponse struct {
	Result SimpleOkResult `json:"result"`
	Error  *struct{}      `json:"error"`
}

type SimpleOkResult struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Result *struct{} `json:"result"`
	Error  string    `json:"error"`
}
