package dto

type ErrorResponse_5xx struct {
	Message            string `json:"message"`
	RequestId_forDebug string `json:"request_id,omitempty"`
	ErrorCode_forDebug int    `json:"code,omitempty"`
}
