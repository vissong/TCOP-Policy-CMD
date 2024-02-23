package main

import (
	"context"
	"os"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"

	"github.com/vissong/TCOP-Policy-CMD/tcop/entity"
)

var cc *TCOP

func TestMain(m *testing.M) {
	SecretID = os.Getenv(SecretIDEnv)
	SecretKey = os.Getenv(SecretKeyEnv)
	credential := common.NewCredential(
		SecretID,
		SecretKey,
	)

	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.Endpoint = "monitor.tencentcloudapi.com"
	client, _ := monitor.NewClient(credential, "ap-guangzhou", clientProfile)

	cc = NewTCOP(client)

	m.Run()
}

func TestTCOP_GetPolicyByID(t *testing.T) {
	r, err := cc.GetPolicyByID(context.Background(), "policy-8e18ig2d")
	if err != nil {
		t.Log(err)
	}
	t.Logf("%+v", r)
}

func TestTCOP_BindResourcesByTag(t1 *testing.T) {
	// r, err := cc.CreateAlarmPolicy(
	// 	context.Background(), &CreateAlarmParams{
	// 		Name:                "test1",
	// 		Remark:              "",
	// 		Namespace:           "redis_mem_edition",
	// 		ConditionTemplateId: 8362782,
	// 		NoticeIDs:           nil,
	// 		Tags:                nil,
	// 	},
	// )
	// if err != nil {
	// 	t1.Error(err)
	// }
	err := cc.BindResourcesByTag(
		context.Background(), "policy-gbn0qqb4", []entity.Tag{
			{
				Key:   "用途",
				Value: "测试",
			},
		},
	)
	if err != nil {
		t1.Error(err)
	}
}
