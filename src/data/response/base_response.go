package response

// BaseResponse represents the base structure of the API response
// @Description Base response structure
type BaseResponse struct {
	Code    int         `json:"code" extensions:"x-order=0"`    // HTTP status code of the response
	Status  string      `json:"status" extensions:"x-order=1"`  // Status of the response
	Message string      `json:"message" extensions:"x-order=2"` // Message of the response
	Data    interface{} `json:"data" extensions:"x-order=3"`    // Data payload of the response
}
