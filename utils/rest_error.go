package errors

type RestErr struct {
	Message string `json: "Message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}
