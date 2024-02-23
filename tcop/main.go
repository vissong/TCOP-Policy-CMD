package main

import (
	"context"
	"os"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"gopkg.in/yaml.v3"
)

const (
	SecretIDEnv  = "MyCloudSecretId"
	SecretKeyEnv = "MyCloudSecretKey"
)

var (
	SecretID  string
	SecretKey string
)

var demoConfig = `
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

	c := &Config{}
	err := yaml.Unmarshal([]byte(demoConfig), c)
	if err != nil {
		panic(err)
	}
	err = SetupAlarmPolicy(ctx, t, c)
	if err != nil {
		panic(err)
	}
}
