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
	str := "{\"header\":{\"namespace\":\"DuerOS.ConnectedHome.Control\",\"name\":\"TurnOffRequest\",\"payLoadVersion\":\"1\",\"messageId\":\"ecc9aec5125b406c8f9735ed524b3ea8_DCS-10-54-221-29-2039-1215161009-973649_0#2_0_Smarthome_657c09e249f607.99938971_6ff647b8\"},\"payload\":{\"accessToken\":\"dc660515-9a1e-11ee-95dd-ce47403e70e1\",\"appliance\":{\"applianceTypes\":null,\"applianceId\":\"0202A0000001-3\",\"modelName\":\"\",\"version\":\"\",\"friendlyName\":\"\",\"friendlyDescription\":\"\",\"isReachable\":false,\"actions\":null,\"additionalApplianceDetails\":{\"deviceType\":\"4-002\",\"deviceSubTypeNo\":\"4-002-001\"},\"manufacturerName\":\"\",\"attributes\":null,\"subType\":\"\"},\"detalValue\":null,\"deltValue\":null,\"targetTemperature\":null,\"fanSpeed\":null,\"lockState\":\"\"}}"
	var protocol model.Protocol
	err = json.Unmarshal([]byte(str), &protocol)
	if err != nil {
		return
	}
	fmt.Printf("%+v", protocol)
	marshal, err := json.Marshal(protocol)
	if err != nil {
		return
	}
	for i := 0; i < 30; i++ {
		err = client.Publish(myconfig.GConfig.NsqdConfig.Topic, marshal)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
	fmt.Println("end======================================")
	fmt.Println("end======================================")
	fmt.Println("end======================================")
	fmt.Println("end======================================")
	fmt.Println("end======================================")
}

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
