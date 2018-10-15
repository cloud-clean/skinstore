package mqttClient

import (
	"encoding/json"
	"skinstore/Entity/lot"
)

var Mc *MqttClient

func MqttInit(){
	Mc = NewMqttClient("lot_server","Kiw28&4292si","lot")
}


func SendMsg(msg lot.LampStatusEntity){
	b,_ := json.Marshal(msg)
	Mc.Publish(string(b))
}
