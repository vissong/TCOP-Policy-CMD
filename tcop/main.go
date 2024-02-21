package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

const (
	SecretIDEnv  = "MyCloudSecretId"
	SecretKeyEnv = "MyCloudSecretKey"
)

var (
	SecretID  string
	SecretKey string
)

func main() {
	SecretID = os.Getenv(SecretIDEnv)
	SecretKey = os.Getenv(SecretKeyEnv)
	credential := common.NewCredential(
		SecretID,
		SecretKey,
	)

	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.Endpoint = "monitor.tencentcloudapi.com"
	client, _ := monitor.NewClient(credential, "ap-guangzhou", clientProfile)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t := NewTCOP(client)

	r, err := t.ListAlarmPolicy(ctx, 1, 20)
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	for _, policy := range r {
		fmt.Printf("%+v", policy)
	}
}
