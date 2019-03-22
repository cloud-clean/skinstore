package mqttClient

import (
	"encoding/json"
	"skinstore/Entity/lot"
)

var Mc *MqttClient

func MqttInit(){
	Mc = NewMqttClient("lot","cloud","root/lot")
}


func SendMsg(msg lot.LampStatusEntity){
	b,_ := json.Marshal(msg)
	Mc.Publish(string(b))
}
