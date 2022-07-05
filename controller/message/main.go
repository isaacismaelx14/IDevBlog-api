package message_ctrl

type IError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type IMessage struct {
	Error *IError     `json:"error"`
	Data  interface{} `json:"data"`
}

func Message(data interface{}) IMessage {
	return IMessage{
		Error: nil,
		Data:  data,
	}
}

func Error(code int, message string) IMessage {
	return IMessage{
		Error: &IError{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}
}
