package mqttClient


var Mc *MqttClient

func MqttInit(){
	Mc = NewMqttClient("admin","cloudhai","test")
}
