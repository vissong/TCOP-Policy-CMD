package main

import (
	"github.com/vissong/TCOP-Policy-CMD/tcop/entity"
)

type Config struct {
	Policies     []Policy      `yaml:"policies"`
	ResourceTags []ResourceTag `yaml:"resourceTags"`
}

type Policy struct {
	Namespace   string       `yaml:"namespace"`
	ConditionID int          `yaml:"conditionID"`
	NoticeIDs   []string     `yaml:"noticeIDs"`
	Name        string       `yaml:"name"`
	Remark      string       `yaml:"remark"`
	Tags        []entity.Tag `yaml:"tags"`
}

type ResourceTag struct {
	Key    string   `yaml:"key"`
	Values []string `yaml:"values"`
}
