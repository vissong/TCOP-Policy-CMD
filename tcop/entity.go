package main

import (
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
)

// Policy 策略对象，从腾讯云对象简化，只保留有用的字段
type Policy struct {
	// 告警策略 ID
	PolicyId string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警策略名称
	PolicyName string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 备注信息
	Remark string `json:"Remark,omitempty" name:"Remark"`
	// 监控类型 MT_QCE=云产品监控
	MonitorType string `json:"MonitorType,omitempty" name:"MonitorType"`

	// 启停状态 0=停用 1=启用
	Enable int64 `json:"Enable,omitempty" name:"Enable"`

	// 策略组绑定的实例数
	UseSum int64 `json:"UseSum,omitempty" name:"UseSum"`

	// 告警策略类型
	Namespace string `json:"Namespace,omitempty" name:"Namespace"`

	// 触发条件模板 Id
	ConditionTemplateId string `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`

	// 指标触发条件
	Condition monitor.AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`

	// 事件触发条件
	EventCondition monitor.AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`

	// 通知规则 id 列表
	NoticeIds []string `json:"NoticeIds,omitempty" name:"NoticeIds"`

	// 通知规则 列表，显示的是 NoticeIds 里的内容，暂时不需要
	// Notices []monitor.AlarmNotice `json:"Notices,omitempty" name:"Notices"`

	// 最后编辑的用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastEditUin string `json:"LastEditUin,omitempty" name:"LastEditUin"`

	// 更新时间
	UpdateTime int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间
	InsertTime int64 `json:"InsertTime,omitempty" name:"InsertTime"`

	// 地域
	Region []string `json:"Region,omitempty" name:"Region"`

	// namespace 显示名字
	NamespaceShowName string `json:"NamespaceShowName,omitempty" name:"NamespaceShowName"`

	// 实例分组ID
	InstanceGroupId int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`

	// 实例分组总实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceSum int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`

	// 实例分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceGroupName string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`

	// 触发条件类型 STATIC=静态阈值 DYNAMIC=动态类型
	RuleType string `json:"RuleType,omitempty" name:"RuleType"`

	// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID
	OriginId string `json:"OriginId,omitempty" name:"OriginId"`

	// 标签
	TagInstances []TagInstance `json:"TagInstances,omitempty" name:"TagInstances"`

	// 过滤条件
	Filter monitor.AlarmConditionFilter `json:"Filter,omitempty" name:"Filter"`

	// 聚合条件
	GroupBy []AlarmGroupByItem `json:"GroupBy,omitempty" name:"GroupBy"`

	// 策略关联的过滤维度信息
	FilterDimensionsParam string `json:"FilterDimensionsParam,omitempty" name:"FilterDimensionsParam"`

	// 是否为一键告警策略
	IsOneClick int64 `json:"IsOneClick,omitempty" name:"IsOneClick"`

	// 一键告警策略是否开启
	OneClickStatus int64 `json:"OneClickStatus,omitempty" name:"OneClickStatus"`

	// 策略是否是全部对象策略
	IsBindAll int64 `json:"IsBindAll,omitempty" name:"IsBindAll"`

	// 策略标签
	Tags []PolicyTag `json:"Tags,omitempty" name:"Tags"`

	// 是否支持告警标签
	IsSupportAlarmTag int64 `json:"IsSupportAlarmTag,omitempty" name:"IsSupportAlarmTag"`
}

// PolicyTag 告警策略标签
type PolicyTag struct {
	Key   string
	Value string
}

type TagInstance struct {
	// 标签Key
	Key string `json:"Key,omitempty" name:"Key"`
	// 标签Value
	Value string `json:"Value,omitempty" name:"Value"`
	// 实例个数
	InstanceSum int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 产品类型，如：cvm
	ServiceType string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 地域ID
	RegionId int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 绑定状态，2：绑定成功，1：绑定中
	BindingStatus int64 `json:"BindingStatus,omitempty" name:"BindingStatus"`
	// 标签状态，2：标签存在，1：标签不存在
	TagStatus int64 `json:"TagStatus,omitempty" name:"TagStatus"`
}

type AlarmGroupByItem struct {
	// Item Id
	Id string `json:"Id,omitempty" name:"Id"`
	// 名称
	Name string `json:"Name,omitempty" name:"Name"`
}
