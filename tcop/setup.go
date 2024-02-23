package main

import (
	"context"
	"log"

	"github.com/vissong/TCOP-Policy-CMD/tcop/entity"
)

func SetupAlarmPolicy(ctx context.Context, t *TCOP, config *Config) error {
	for _, policy := range config.Policies {
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
			if err := t.BindResourcesByTag(ctx, p.PolicyId, convertTags(config.ResourceTags)); err != nil {
				continue
			}
		} else {
			// 更新逻辑，只更新搜到的第一个策略
			if err := t.BindResourcesByTag(
				ctx, existPolicies.Policies[0].PolicyId,
				convertTags(config.ResourceTags),
			); err != nil {
				continue
			}
		}
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
