package main

import (
	"context"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

type TCOP struct {
	client *monitor.Client
}

func NewTCOP(client *monitor.Client) *TCOP {
	return &TCOP{client: client}
}

func (t *TCOP) ListAlarmPolicy(ctx context.Context, pageNum,
	pageSize int64) ([]*Policy, error) {

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
	var ret []*Policy
	for _, policy := range response.Response.Policies {
		ret = append(ret, convertPolicy(policy))
	}
	return ret, nil
}

// convertPolicy 对象转换，腾讯云的对象字段都是指针，十分的不友好
func convertPolicy(policy *monitor.AlarmPolicy) *Policy {
	r := &Policy{
		PolicyId:            *policy.PolicyId,
		PolicyName:          *policy.PolicyName,
		Remark:              *policy.Remark,
		MonitorType:         *policy.MonitorType,
		Enable:              *policy.Enable,
		UseSum:              *policy.UseSum,
		Namespace:           *policy.Namespace,
		ConditionTemplateId: *policy.ConditionTemplateId,
		Condition:           *policy.Condition,
		EventCondition:      *policy.EventCondition,
		LastEditUin:         *policy.LastEditUin,
		UpdateTime:          *policy.UpdateTime,
		InsertTime:          *policy.InsertTime,
		// Region:                *policy.Region,
		NamespaceShowName: *policy.NamespaceShowName,
		InstanceGroupId:   *policy.InstanceGroupId,
		InstanceSum:       *policy.InstanceSum,
		InstanceGroupName: *policy.InstanceGroupName,
		RuleType:          *policy.RuleType,
		OriginId:          *policy.OriginId,
		// TagInstances:          *policy.TagInstances,
		Filter: monitor.AlarmConditionFilter{},
		// GroupBy:               *policy.GroupBy,
		FilterDimensionsParam: *policy.FilterDimensionsParam,
		IsOneClick:            *policy.IsOneClick,
		OneClickStatus:        *policy.OneClickStatus,
		IsBindAll:             *policy.IsBindAll,
		IsSupportAlarmTag:     *policy.IsSupportAlarmTag,
		Tags:                  []PolicyTag{},
		NoticeIds:             []string{},
	}
	for _, id := range policy.NoticeIds {
		r.NoticeIds = append(r.NoticeIds, *id)
	}
	for _, tag := range policy.Tags {
		r.Tags = append(
			r.Tags, PolicyTag{
				Key:   *tag.Key,
				Value: *tag.Value,
			},
		)
	}
	return r
}
