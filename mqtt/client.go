package mqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

const(
	MQTT_SERVER = "tcp://202.182.118.148:61613"
	CLIENT_ID = "sample"
)
type MqttClient struct {
	Client mqtt.Client
	Topic string
}

func NewMqttClient(username,password,topic string) *MqttClient{
	opts := mqtt.NewClientOptions().AddBroker(MQTT_SERVER).SetClientID(CLIENT_ID)
	opts.SetProtocolVersion(4)
	opts.SetUsername(username)
	opts.SetPassword(password)
	c := mqtt.NewClient(opts)
	if token:= c.Connect();token.Wait() && token.Error() != nil{
		fmt.Println(token.Error())
		panic(token.Error())
	}
	return &MqttClient{Client:c,Topic:topic}
}

func (self *MqttClient) Publish(msg string) bool{
	if token := self.Client.Publish(self.Topic,1,false,msg);token.Wait()&&token.Error() != nil{
		return false
	}else{
		return true
	}
}

func (self *MqttClient) Subscribe(handler mqtt.MessageHandler){
	if token:=self.Client.Subscribe(self.Topic,0,handler);token.Wait() && token.Error() != nil{

	}
}