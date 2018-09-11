package mqttClient


var Mc *MqttClient

func MqttInit(){
	Mc = NewMqttClient("lot_server","Kiw28&4292si","lot")
}
