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
	AlarmModule = "monitor"
	MonitorType = "MT_QCE" // 云产品监控类型
)

type TCOP struct {
	client *monitor.Client
	config *Config
}

// NewTCOP 创建 tcop 实例
func NewTCOP(client *monitor.Client, config *Config) *TCOP {
	return &TCOP{
		client: client,
		config: config,
	}
}

// SearchAlarmPolicyByName 基于监控策略名字搜索
func (t *TCOP) SearchAlarmPolicyByName(ctx context.Context, nameKeyword string, matchAll bool) (*entity.AlarmPolicies,
	error) {
	log.Printf("nameKeyword is %s", nameKeyword)
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
			// 名称不完全匹配
			if !matchAll && strings.Contains(item.PolicyName, nameKeyword) {
				result.Policies = append(result.Policies, item)
			}
			// 名称完全匹配
			if matchAll && item.PolicyName == nameKeyword {
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
	request.Module = common.StringPtr(AlarmModule)
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
	ConditionTemplateId int64
	NoticeIDs           []string
	Tags                []entity.Tag
	ProjectId           int64
}

// CreateAlarmPolicy 创建告警接口，由于接口不支持 IsBindAll 所以只支持按照标签绑定的方式进行调整告警
func (t *TCOP) CreateAlarmPolicy(ctx context.Context, params *CreateAlarmParams) (*entity.CreatePolicyResult, error) {
	if params.ConditionTemplateId == 0 {
		return nil, fmt.Errorf("createAlarmParams.ConditionTemplateId is Requreid")
	}
	request := monitor.NewCreateAlarmPolicyRequest()
	request.Module = common.StringPtr(AlarmModule)
	request.MonitorType = common.StringPtr(MonitorType)
	request.PolicyName = common.StringPtr(params.Name)
	request.Remark = common.StringPtr(params.Remark)
	request.Namespace = common.StringPtr(params.Namespace)
	request.ConditionTemplateId = common.Int64Ptr(params.ConditionTemplateId)
	request.ProjectId = common.Int64Ptr(params.ProjectId)
	for _, d := range params.NoticeIDs {
		request.NoticeIds = append(request.NoticeIds, common.StringPtr(d))
	}
	for _, i := range params.Tags {
		request.Tags = append(
			request.Tags, &monitor.Tag{
				Key:   common.StringPtr(i.Key),
				Value: common.StringPtr(i.Value),
			},
		)
	}
	response, err := t.client.CreateAlarmPolicyWithContext(ctx, request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("CreateAlarmPolicyWithContext An API error has returned: %s", err)
	}
	if err != nil {
		return nil, err
	}

	return convert(
		response.ToJsonString(), "Response",
		&entity.CreatePolicyResult{},
	).(*entity.CreatePolicyResult), nil
}

// GetPolicyByID 获取单个告警策略
func (t *TCOP) GetPolicyByID(ctx context.Context, policyID string) (*entity.Policy, error) {
	req := monitor.NewDescribeAlarmPolicyRequest()
	req.Module = common.StringPtr(AlarmModule)
	req.PolicyId = common.StringPtr(policyID)
	resp, err := t.client.DescribeAlarmPolicyWithContext(ctx, req)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
	}
	if err != nil {
		return nil, err
	}

	return convert(resp.ToJsonString(), "Response.Policy", &entity.Policy{}).(*entity.Policy), nil
}

// BindResourcesByTag 绑定资源 tag 到具体的告警策略上
func (t *TCOP) BindResourcesByTag(ctx context.Context, policyID string, tags []entity.Tag) error {
	policy, err := t.GetPolicyByID(ctx, policyID)
	if err != nil {
		return err
	}
	req := monitor.NewBindingPolicyTagRequest()
	req.Module = common.StringPtr(AlarmModule)
	req.PolicyId = common.StringPtr(policyID)
	req.GroupId = common.StringPtr("0")
	req.ServiceType = common.StringPtr(policy.Namespace)
	for _, tag := range tags {
		req.BatchTag = append(
			req.BatchTag, &monitor.PolicyTag{
				Key:   common.StringPtr(tag.Key),
				Value: common.StringPtr(tag.Value),
			},
		)
	}
	// 该请求没有有意义的回包内容
	_, err = t.client.BindingPolicyTagWithContext(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// ListConditionTemplate 获取触发条件模板, id 和 name 分别传一个即可
func (t *TCOP) ListConditionTemplate(ctx context.Context, id string, name string, namespace string,
	num int64) (*entity.ConditionsTemplateList, error) {
	req := monitor.NewDescribeConditionsTemplateListRequest()
	req.Module = common.StringPtr(AlarmModule)
	req.ViewName = common.StringPtr(namespace)
	if id != "" {
		req.GroupID = common.StringPtr(id)
	}
	if name != "" {
		req.GroupName = common.StringPtr(name)
	}
	req.Limit = common.Int64Ptr(num)

	resp, err := t.client.DescribeConditionsTemplateListWithContext(ctx, req)
	if err != nil {
		return nil, err
	}
	return convert(
		resp.ToJsonString(),
		"Response", &entity.ConditionsTemplateList{},
	).(*entity.ConditionsTemplateList), nil
}

// 从 json 中得到数据之后，解析为目标对象
func convert(inputJSON string, path string, entity interface{}) interface{} {
	body := gjson.Get(inputJSON, path).String()
	// log.Println(body)
	_ = json.Unmarshal([]byte(body), entity)
	return entity
}
