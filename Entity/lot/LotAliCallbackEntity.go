package lot


type AliCallback struct {
	Header struct{
		MessageId string   		`json:"messageId"`
		Name string				`json:"name"`
		Namespace string		`json:"namespace"`
		PayLoadVersion int		`json:"payLoadVersion"`
	} 							`json:"header"`
	Payload struct{
		accessToken string
	}							`json:"payload"`

}
