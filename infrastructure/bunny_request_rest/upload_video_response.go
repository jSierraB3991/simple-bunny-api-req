package bunnyrequestrest

type UploadVideoResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
