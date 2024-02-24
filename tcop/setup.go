package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vissong/TCOP-Policy-CMD/tcop/entity"
)

// PolicyResult 策略初始化之后的结果
type PolicyResult struct {
	Name string
	ID   string
	URL  string
}

const (
	urlFormat = "https://console.cloud.tencent.com/monitor/alarm/policy/detail/%s"
)

func SetupAlarmPolicy(ctx context.Context, t *TCOP, config *Config) error {
	var result []PolicyResult
	for _, policy := range config.Policies {
		var _policyID string
		log.Printf("处理告警策略: %s \n", policy.Name)
		existPolicies, err := t.SearchAlarmPolicyByName(ctx, policy.Name)
		if err != nil {
			continue
		}
		// 创建逻辑
		if existPolicies.TotalCount == 0 {
			p, err := t.CreateAlarmPolicy(
				ctx, &CreateAlarmParams{
					Name:                policy.Name,
					Remark:              policy.Remark,
					Namespace:           policy.Namespace,
					ConditionTemplateId: int64(policy.ConditionID),
					NoticeIDs:           policy.NoticeIDs,
					Tags:                policy.Tags,
				},
			)
			if err != nil {
				continue
			}
			log.Printf("policy create success, policy id is %s", p.PolicyId)
			if err := t.BindResourcesByTag(ctx, p.PolicyId, convertTags(config.ResourceTags)); err != nil {
				continue
			}
			_policyID = p.PolicyId
		} else {
			// 更新逻辑，只更新搜到的第一个策略
			if err := t.BindResourcesByTag(
				ctx, existPolicies.Policies[0].PolicyId,
				convertTags(config.ResourceTags),
			); err != nil {
				continue
			}
			_policyID = existPolicies.Policies[0].PolicyId
		}
		log.Printf(
			"bind resoureces to policy %s by tag success, tags is %+v",
			_policyID,
			config.ResourceTags,
		)
		result = append(
			result, PolicyResult{
				Name: policy.Name,
				ID:   _policyID,
				URL:  fmt.Sprintf(urlFormat, _policyID),
			},
		)
	}

	log.Printf("Policy List Is:")
	for _, policyResult := range result {
		log.Println("------")
		log.Printf("Name: %s", policyResult.Name)
		log.Printf("LINK: %s", policyResult.URL)
	}
	return nil
}

func convertTags(inputTags []ResourceTag) []entity.Tag {
	var result []entity.Tag
	for _, inputTag := range inputTags {
		for _, value := range inputTag.Values {
			result = append(
				result, entity.Tag{
					Key:   inputTag.Key,
					Value: value,
				},
			)
		}
	}
	return result
}
