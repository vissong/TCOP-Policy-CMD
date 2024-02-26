package entity

// AlarmPolicies 通过 DescribeAlarmPolicies 接口获得的回包，描述完整的监控策略
type AlarmPolicies struct {
	TotalCount int      `json:"TotalCount"`
	Policies   []Policy `json:"Policies"`
	RequestId  string   `json:"RequestId"`
}

// Policy 告警策略对象
type Policy struct {
	PolicyId            string `json:"PolicyId"`
	PolicyName          string `json:"PolicyName"`
	Remark              string `json:"Remark"`
	MonitorType         string `json:"MonitorType"`
	Enable              int    `json:"Enable"`
	UseSum              int    `json:"UseSum"`
	ProjectId           int    `json:"ProjectId"`
	ProjectName         string `json:"ProjectName"`
	Namespace           string `json:"Namespace"`
	ConditionTemplateId string `json:"ConditionTemplateId"`
	Condition           struct {
		IsUnionRule int `json:"IsUnionRule"`
		Rules       []struct {
			MetricName      string `json:"MetricName"`
			Period          int    `json:"Period"`
			Operator        string `json:"Operator"`
			Value           string `json:"Value"`
			ContinuePeriod  int    `json:"ContinuePeriod"`
			NoticeFrequency int    `json:"NoticeFrequency"`
			IsPowerNotice   int    `json:"IsPowerNotice"`
			Filter          struct {
				Type       string `json:"Type"`
				Dimensions string `json:"Dimensions"`
			} `json:"Filter"`
			Description       string `json:"Description"`
			Unit              string `json:"Unit"`
			RuleType          string `json:"RuleType"`
			IsAdvanced        int    `json:"IsAdvanced"`
			IsOpen            int    `json:"IsOpen"`
			ProductId         string `json:"ProductId"`
			ValueMax          int    `json:"ValueMax"`
			ValueMin          int    `json:"ValueMin"`
			HierarchicalValue struct {
				Remind  string `json:"Remind"`
				Warn    string `json:"Warn"`
				Serious string `json:"Serious"`
			} `json:"HierarchicalValue"`
		} `json:"Rules"`
		ComplexExpression string `json:"ComplexExpression"`
	} `json:"Condition"`
	EventCondition struct {
		Rules []struct {
			MetricName      string `json:"MetricName"`
			Period          int    `json:"Period"`
			Operator        string `json:"Operator"`
			Value           string `json:"Value"`
			ContinuePeriod  int    `json:"ContinuePeriod"`
			NoticeFrequency int    `json:"NoticeFrequency"`
			IsPowerNotice   int    `json:"IsPowerNotice"`
			Description     string `json:"Description"`
			Unit            string `json:"Unit"`
			RuleType        string `json:"RuleType"`
			IsAdvanced      int    `json:"IsAdvanced"`
			IsOpen          int    `json:"IsOpen"`
			ProductId       string `json:"ProductId"`
		} `json:"Rules"`
	} `json:"EventCondition"`
	NoticeIds []string `json:"NoticeIds"`
	Notices   []struct {
		Id          string `json:"Id"`
		Name        string `json:"Name"`
		UpdatedAt   string `json:"UpdatedAt"`
		UpdatedBy   string `json:"UpdatedBy"`
		NoticeType  string `json:"NoticeType"`
		UserNotices []struct {
			ReceiverType          string        `json:"ReceiverType"`
			StartTime             int           `json:"StartTime"`
			EndTime               int           `json:"EndTime"`
			NoticeWay             []string      `json:"NoticeWay"`
			UserIds               []int         `json:"UserIds"`
			GroupIds              []interface{} `json:"GroupIds"`
			PhoneOrder            []interface{} `json:"PhoneOrder"`
			PhoneCircleTimes      int           `json:"PhoneCircleTimes"`
			PhoneInnerInterval    int           `json:"PhoneInnerInterval"`
			PhoneCircleInterval   int           `json:"PhoneCircleInterval"`
			NeedPhoneArriveNotice int           `json:"NeedPhoneArriveNotice"`
			Weekday               []int         `json:"Weekday"`
			OnCallFormIDs         []string      `json:"OnCallFormIDs"`
		} `json:"UserNotices"`
		URLNotices     []interface{} `json:"URLNotices"`
		IsPreset       int           `json:"IsPreset"`
		NoticeLanguage string        `json:"NoticeLanguage"`
		AMPConsumerId  string        `json:"AMPConsumerId"`
		CLSNotices     []interface{} `json:"CLSNotices"`
	} `json:"Notices"`
	TriggerTasks   []interface{} `json:"TriggerTasks"`
	ConditionsTemp struct {
		TemplateName string `json:"TemplateName"`
	} `json:"ConditionsTemp"`
	LastEditUin           string        `json:"LastEditUin"`
	UpdateTime            int           `json:"UpdateTime"`
	InsertTime            int           `json:"InsertTime"`
	Region                []string      `json:"Region"`
	NamespaceShowName     string        `json:"NamespaceShowName"`
	IsDefault             int           `json:"IsDefault"`
	CanSetDefault         int           `json:"CanSetDefault"`
	InstanceGroupId       int           `json:"InstanceGroupId"`
	InstanceSum           int           `json:"InstanceSum"`
	InstanceGroupName     string        `json:"InstanceGroupName"`
	RuleType              string        `json:"RuleType"`
	OriginId              string        `json:"OriginId"`
	TagInstances          []interface{} `json:"TagInstances"`
	FilterDimensionsParam string        `json:"FilterDimensionsParam"`
	IsOneClick            int           `json:"IsOneClick"`
	OneClickStatus        int           `json:"OneClickStatus"`
	AdvancedMetricNumber  int           `json:"AdvancedMetricNumber"`
	IsBindAll             int           `json:"IsBindAll"`
	Tags                  []Tag         `json:"Tags"`
	IsSupportAlarmTag     int           `json:"IsSupportAlarmTag"`
}

// Tag 标签
type Tag struct {
	Key   string `json:"Key" yaml:"key"`
	Value string `json:"Value" yaml:"value"`
}

// CreatePolicyResult 创建告警策略的结果
type CreatePolicyResult struct {
	PolicyId  string `json:"PolicyId"`
	OriginId  string `json:"OriginId"`
	RequestId string `json:"RequestId"`
}

type ConditionsTemplateList struct {
	Total             int `json:"Total"`
	TemplateGroupList []struct {
		Conditions      []Condition   `json:"Conditions"`
		EventConditions []interface{} `json:"EventConditions"`
		PolicyGroups    []struct {
			CanSetDefault           bool   `json:"CanSetDefault"`
			GroupID                 int    `json:"GroupID"`
			GroupName               string `json:"GroupName"`
			InsertTime              int    `json:"InsertTime"`
			IsDefault               int    `json:"IsDefault"`
			Enable                  bool   `json:"Enable"`
			LastEditUin             int    `json:"LastEditUin"`
			NoShieldedInstanceCount int    `json:"NoShieldedInstanceCount"`
			ParentGroupID           int    `json:"ParentGroupID"`
			ProjectID               int    `json:"ProjectID"`
			Remark                  string `json:"Remark"`
			UpdateTime              int    `json:"UpdateTime"`
			TotalInstanceCount      int    `json:"TotalInstanceCount"`
			ViewName                string `json:"ViewName"`
			IsUnionRule             int    `json:"IsUnionRule"`
		} `json:"PolicyGroups"`
		GroupID     int    `json:"GroupID"`
		GroupName   string `json:"GroupName"`
		InsertTime  int    `json:"InsertTime"`
		LastEditUin int    `json:"LastEditUin"`
		Remark      string `json:"Remark"`
		UpdateTime  int    `json:"UpdateTime"`
		ViewName    string `json:"ViewName"`
		IsUnionRule int    `json:"IsUnionRule"`
	} `json:"TemplateGroupList"`
	RequestId string `json:"RequestId"`
}

type Condition struct {
	AlarmNotifyPeriod int    `json:"AlarmNotifyPeriod"`
	AlarmNotifyType   int    `json:"AlarmNotifyType"`
	CalcType          string `json:"CalcType"`
	CalcValue         string `json:"CalcValue"`
	ContinueTime      string `json:"ContinueTime"`
	MetricID          int    `json:"MetricID"`
	MetricDisplayName string `json:"MetricDisplayName"`
	Period            int    `json:"Period"`
	RuleID            int    `json:"RuleID"`
	Unit              string `json:"Unit"`
	IsAdvanced        int    `json:"IsAdvanced"`
	IsOpen            int    `json:"IsOpen"`
	ProductId         string `json:"ProductId"`
}
