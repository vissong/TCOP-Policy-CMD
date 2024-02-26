package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"gopkg.in/yaml.v3"
)

var (
	SecretID   string
	SecretKey  string
	Action     string
	ConfigFile string
)

func main() {
	flag.StringVar(&Action, "action", "set", "cmd action type, now only support [set] action(default)")
	flag.StringVar(&ConfigFile, "config", "", "yaml config path")
	flag.Parse()
	log.Printf("config path is %s", ConfigFile)
	log.Printf("action is %s", Action)
	if ConfigFile == "" {
		log.Printf("params [config] is required, please check it")
		return
	}

	content, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Printf("read config file failed, err is %s", err)
		return
	}
	c := &Config{}
	if err := yaml.Unmarshal(content, c); err != nil {
		log.Printf("parse config failed, err is %s", err)
		return
	}

	tcop, err := initClient(c)
	if err != nil {
		log.Printf("init TCOP client failed, err is %s", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if Action == "set" {
		if err := SetupAlarmPolicy(ctx, tcop, c); err != nil {
			log.Printf("setup alarm policy failed, err is %s", err)
		}
	}
}

func initClient(config *Config) (*TCOP, error) {
	if config.Secret.LoadFromEnv {
		SecretID = os.Getenv(config.Secret.SecretID)
		SecretKey = os.Getenv(config.Secret.SecretKey)
	} else {
		SecretID = config.Secret.SecretID
		SecretKey = config.Secret.SecretKey
	}
	credential := common.NewCredential(
		SecretID,
		SecretKey,
	)
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.Endpoint = "monitor.tencentcloudapi.com"
	client, err := monitor.NewClient(credential, "ap-guangzhou", clientProfile)
	if err != nil {
		return nil, err
	}
	return NewTCOP(client, config), nil
}
