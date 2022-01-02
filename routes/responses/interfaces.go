package responses

type SuccessResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

type AccountCreateResponse struct {
	Success bool `json:"success"`
	Token string `json:"token"`
}
