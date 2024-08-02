package response

// BaseResponse represents the base structure of the API response
// @Description Base response structure
type BaseResponse struct {
	Code    int         `json:"code" example:"200" extensions:"x-order=0"`                                 // HTTP status code of the response
	Status  string      `json:"status" example:"success" extensions:"x-order=1"`                           // Status of the response
	Message string      `json:"message" example:"Operation completed successfully" extensions:"x-order=2"` // Message of the response
	Data    interface{} `json:"data" extensions:"x-order=3"`                                               // Data payload of the response
}
