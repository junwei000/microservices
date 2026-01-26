package request

type SendSms struct {
	Phone string `json:"phone" binding:"required"`
}
