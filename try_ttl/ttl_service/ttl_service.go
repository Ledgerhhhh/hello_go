package ttl_service

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"sync"
)

type TTSService struct {
	initOnce  sync.Once
	ttsClient *sdk.Client
}

type TokenResult struct {
	ErrMsg string
	Token  struct {
		UserId     string
		Id         string
		ExpireTime int64
	}
}

func (receiver *TTSService) CreateToken() *TokenResult {
	client, err := sdk.NewClientWithAccessKey("cn-shanghai", "LTAI5tFj1BR4xH9nmgrniZYd", "xnWxmsA5vvveSB9JE0rZSZR3elftKU")
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
	request.ApiName = "CreateToken"
	request.Version = "2019-02-28"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpStatus())
	fmt.Print(response.GetHttpContentString())

	var tr TokenResult
	err = json.Unmarshal([]byte(response.GetHttpContentString()), &tr)
	if err == nil {
		fmt.Println(tr.Token.Id)
		fmt.Println(tr.Token.ExpireTime)
	} else {
		fmt.Println(err)
	}
	return &tr
}
