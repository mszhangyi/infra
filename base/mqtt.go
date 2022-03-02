package base

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mszhangyi/infra"
	"time"
)

const (
	PubClientSize = 3
)

var (
	pubChanMap map[int]chan *ResponseMsg
	mClient    mqtt.Client
	roundRound int
)

type ResponseMsg struct {
	Topic string
	Msg   interface{}
}

type MQttStarter struct {
	infra.BaseStarter
}

func (t *MQttStarter) Init() {
	pubChanMap = make(map[int]chan *ResponseMsg)
	for i := 0; i < PubClientSize; i++ {
		ch := make(chan *ResponseMsg, 1000)
		pubChanMap[i] = ch
		go t.startPublishMQtt(i)
	}
}
func (t *MQttStarter) Stop() {
	//关闭服务的时候
}

/**
发送消息
*/
func SendMsg(pub string, msg string) {
	response := &ResponseMsg{
		Topic: pub,
		Msg:   msg,
	}
	if roundRound >= PubClientSize {
		roundRound = 0
	}
	pubChanMap[roundRound%PubClientSize] <- response
	roundRound++
}

//---------------------------------------------------------------------------------   启动发布Topic
func (t *MQttStarter) startPublishMQtt(index int) {
	clientId := fmt.Sprintf("%s-pub-%d-%d", props.Name, index, time.Now().UnixNano())
	opts := getMqttOpts(clientId)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	for {
		select {
		case t, ok := <-pubChanMap[index]:
			if !ok {
				return
			}
			c.Publish(t.Topic, 0, false, t.Msg)
		}
	}
	c.Disconnect(3000)
}

func getMqttOpts(ClientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(props.EmqAddr).SetClientID(ClientId)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetConnectTimeout(60 * time.Second)
	opts.SetUsername(props.EmqUser)
	opts.SetPassword(props.EmqPwd)
	opts.SetProtocolVersion(4)
	opts.SetAutoReconnect(true)                 //设置自动重新连接
	opts.SetOnConnectHandler(onConnectCallBack) //设置初始连接时和自动重新连接时调用的函数。
	return opts
}

var onConnectCallBack mqtt.OnConnectHandler = func(client mqtt.Client) {
	/*options := client.OptionsReader()
	clientId := options.ClientID()
	logrus.Info("mqtt " +clientId + " client connect success ")*/
}
