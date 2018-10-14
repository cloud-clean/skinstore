package lot


type AliCallback struct {
	Header struct{
		MessageId string   		`json:"messageId"`
		Name string				`json:"name"`
		Namespace string		`json:"namespace"`
		PayLoadVersion int		`json:"payLoadVersion"`
	} 							`json:"header"`
	Payload struct{
		AccessToken string		`json:"accessToken"`
		Devices []LotDevice	`json:"devices"`
		Actions []string		`json:"actions"`
		Extension map[string]string `json:"extension"`
	}							`json:"payload"`

}

type LotDevice struct {
	DeviceId string 	`json:"deviceId"`
	DeviceName string 	`json:"deviceName"`
	DeviceType string 	`json:"deviceType"`
	Zone 		string	`json:"zone"`
	Model		string 	`json:"model"`
	Icon 		string  `json:"icon"`
	Value 		string 	`json:"value"`
	Attribute	string 	`json:"attribute"`
	Properties []map[string]string 	`json:"properties"`
}
