package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"github.com/tidwall/gjson"

	"github.com/vissong/TCOP-Policy-CMD/tcop/entity"
)

const (
	MaxPageSize = 100
)

type TCOP struct {
	client *monitor.Client
}

func NewTCOP(client *monitor.Client) *TCOP {
	return &TCOP{client: client}
}

// SearchAlarmPolicyByName 基于监控策略名字搜索
func (t *TCOP) SearchAlarmPolicyByName(ctx context.Context, keyword string) (*entity.AlarmPolicies, error) {
	log.Printf("keyword is %s", keyword)
	var page int64 = 1
	result := &entity.AlarmPolicies{}
	maxLoop := 1000

	for i := 0; i < maxLoop; i++ {
		list, err := t.ListAlarmPolicy(ctx, page, MaxPageSize)
		// 翻尽了
		if len(list.Policies) == 0 {
			break
		}
		if err != nil {
			return nil, err
		}
		for _, item := range list.Policies {
			if strings.Contains(item.PolicyName, keyword) {
				result.Policies = append(result.Policies, item)
			}
		}
		page++
	}
	result.TotalCount = len(result.Policies)
	return result, nil
}

// ListAlarmPolicy 分页加载监控策略
func (t *TCOP) ListAlarmPolicy(_ context.Context, pageNum,
	pageSize int64) (*entity.AlarmPolicies, error) {

	request := monitor.NewDescribeAlarmPoliciesRequest()
	request.Module = common.StringPtr("monitor")
	request.PageNumber = common.Int64Ptr(pageNum)
	request.PageSize = common.Int64Ptr(pageSize)

	response, err := t.client.DescribeAlarmPolicies(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
	}
	if err != nil {
		return nil, err
	}

	result := convert(response.ToJsonString(), "Response", &entity.AlarmPolicies{}).(*entity.AlarmPolicies)
	return result, nil
}

type CreateAlarmParams struct {
	Name                string
	Remark              string
	Namespace           string
	ConditionTemplateId string
	NoticeIDs           []string
	Tags                []entity.Tag
}

// func (t *TCOP) CreateAlarmPolicy(ctx context.Context, policies *entity.AlarmPolicies) (*entity.AlarmPolicies, error) {
// 	request := monitor.NewCreateAlarmPolicyRequest()
// 	t.client.CreateAlarmPolicyWithContext(ctx, request)
// }

// 从 json 中得到数据之后，解析为目标对象
func convert(inputJSON string, path string, entity interface{}) interface{} {
	body := gjson.Get(inputJSON, path).String()
	// log.Println(body)
	_ = json.Unmarshal([]byte(body), entity)
	return entity
}
