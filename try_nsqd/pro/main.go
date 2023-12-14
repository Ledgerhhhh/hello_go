package main

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_nsqd/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	client, err := GetProduceClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	protocol := model.Protocol{
		Header: model.Header{
			Namespace: "",
			Name:      "",
			Version:   nil,
			MessageID: "",
		},
		Payload: model.Payload{
			AccessToken:              "",
			DeviceID:                 "",
			DeviceIDs:                nil,
			DeviceType:               "",
			Params:                   nil,
			Attribute:                "",
			Value:                    "",
			Devices:                  nil,
			DeviceResponseList:       nil,
			ErrorCode:                "",
			ErrorMsg:                 "",
			DiscoveredAppliances:     nil,
			DiscoveredGroups:         nil,
			Appliance:                nil,
			Function:                 "",
			ColorTemperatureInKelvin: 0,
			DeltaPercentage:          nil,
			Brightness:               nil,
			Attributes:               nil,
			PreviousState:            nil,
			DependentServiceName:     "",
			DetalValue:               nil,
			DeltValue:                nil,
			Color:                    nil,
			AchievedState:            nil,
			Mode:                     nil,
			TargetTemperature:        nil,
			FanSpeed:                 nil,
			LockState:                "",
		},
	}
	marshal, err := json.Marshal(protocol)
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		err = client.Publish(myconfig.GConfig.NsqdConfig.Topic, marshal)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
}

// GetProduceClient 获取一个生产者客户端连接
func GetProduceClient() (*nsq.Producer, error) {
	config := nsq.NewConfig()
	config.AuthSecret = myconfig.GConfig.NsqdConfig.AuthSecret

	producer, err := nsq.NewProducer(myconfig.GConfig.NsqdConfig.Host+
		":"+
		fmt.Sprintf("%d", myconfig.GConfig.NsqdConfig.Port),
		config,
	)
	if err != nil {
		return nil, errors.New("GetProduceClient error: " + err.Error())
	}
	return producer, nil
}
