package main

import (
	"context"
	"os"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"gopkg.in/yaml.v3"

	"github.com/vissong/TCOP-Policy-CMD/tcop/entity"
)

var cc *TCOP

var demoConfig = `
secret:
  # loadFromEnv = true
  # 从环境变量中获取调用腾讯云 api 的 id 和 key，这里声明的是环境变量的 name
  # loadFromEnv = false
  # 从配置文件读取
  # 建议：优先推荐放到环境变量，而不是放到配置文件中，如果放在配置文件中，请保障配置文件不提交到 git 仓库中
  loadFromEnv: true
  secretID: MyCloudSecretId
  secretKey: MyCloudSecretKey
policies:
  - namespace: redis_mem_edition # 云产品类型
    conditionID:  8362782 # 触发条件模板 id
    noticeIDs: 
      - notice-4uks09xe # 通知渠道模板 id
    name: QQ基础 Redis 实例告警 # 告警策略名称，创建时候使用，更新时候也会用这个查询后修改
    remark: 维护 by tcop cmd  # 备注
    tags: # 监控策略的 tag
      - key: 用途
        value: 魔法
resourceTags: # 资源标签
  - key: 用途
    values:
      - 魔法
      - 测试
`

func TestMain(m *testing.M) {
	c := &Config{}
	err := yaml.Unmarshal([]byte(demoConfig), c)
	if err != nil {
		panic(err)
	}
	if c.Secret.LoadFromEnv {
		SecretID = os.Getenv(c.Secret.SecretID)
		SecretKey = os.Getenv(c.Secret.SecretKey)
	} else {
		SecretID = c.Secret.SecretID
		SecretKey = c.Secret.SecretKey
	}

	credential := common.NewCredential(
		SecretID,
		SecretKey,
	)
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.Endpoint = "monitor.tencentcloudapi.com"
	client, _ := monitor.NewClient(credential, "ap-guangzhou", clientProfile)
	cc = NewTCOP(client, c)

	m.Run()
}

func TestSetupAlarmPolicy(t *testing.T) {
	c := &Config{}
	err := yaml.Unmarshal([]byte(demoConfig), c)
	if err != nil {
		panic(err)
	}
	err = SetupAlarmPolicy(context.Background(), cc, c)
	if err != nil {
		panic(err)
	}
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

func TestTCOP_GetConditionByID(t1 *testing.T) {
	r, err := cc.ListConditionTemplate(context.Background(), "8362782", "", "redis_mem_edition", 1)
	if err != nil {
		t1.Error(err)
	}
	t1.Logf("%+v", r)
}
