# TCOP-Policy-CMD
用于管理腾讯云可观测平台告警策略的命令行工具 / a command tool to manage Tencent Cloud Observable Platform alarm policy

#### 为什么要做这个工具：
在公司内部，资源都在统一的大账号下，使用标签来进行不同业务的资源的标识，在告警策略的设置上，不同的业务
可能会有不同的要求，在控制台操作，针对每个云产品，如果需要变更业务标签返回，就需要重复的进行编辑工作多次，
比较浪费操作时间。

基于此工具，在创建不同云产品的告警出发条件之后，可以通过 yaml 配置的方式，进行不同云产品告警策略的声明
以及指定对应的资源标签范围，有需要变更范围的时候，只需要使用工具执行一次即可。

#### ROADMAP：
- 支持拉取告警策略列表 ✅
- 支持搜索告警列表 ✅
- 支持创建告警策略，支持按照标签绑定，支持使用触发模板来设置出发条件 ✅ 
- 支持编辑告警策略，更换触发条件，告警渠道，更改绑定对象的标签 ✅
- 支持基于 yaml 的模板描述，统一刷新告警策略的范围 ✅
- 支持触发条件模板的拉取/搜索 ✅

#### 能力限制：
- 目前云接口不支持使用【全部对象】的方式进行资源与告警策略的绑定，所以推荐使用标签绑定资源
- 目前腾讯云接口不支持触发条件模板的创建，只支持拉取，建议优先在控制台创建好不同产品的告警触发条件
- 目前部分云产品的告警策略（比如 Ckafka 的 ConsumerGroup+Topic 监控）
不支持标签绑定资源之后配置指标监控，这种建议手动到控制台设置。
  目前发现不支持的产品列表：
  - 消息队列类（tdmq 队列订阅，ckafka consumerGroup+Topic 等）

#### 使用方式

在云控制台创建告警条件和通知模板：

1. 针对不同的云产品，创建告警触发条件模板，并记录 id （url 上的数字 detail 路径后的数字）
2. 创建告警通知模板，用于告警的发送，如果不同的云产品想要发送到不同渠道，请创建多个，并记录 id（如：notice-4uks09xe）

执行工具：

1. 创建密钥：https://console.cloud.tencent.com/cam/capi 建议使用子账号方式，方便控制密钥的权限范围
2. 保存密钥，创建后腾讯云官网不会保存密钥显示，请妥善保存密钥
3. 选择密钥传递方式，推荐放到环境变量中传递给工具
4. 编辑配置文件，完善相关信息
5. 执行工具，如：

```bash
tcop --config ./demo.yaml
```

#### config demo：

配置文件是 yaml 格式的。

```yaml
secret:
  # loadFromEnv = true
  # 从环境变量中获取调用腾讯云 api 的 id 和 key，这里声明的是环境变量的 name
  # loadFromEnv = false
  # 从配置文件读取
  # 建议：优先推荐放到环境变量，而不是放到配置文件中，如果放在配置文件中，请保障配置文件不提交到 git 仓库中
  loadFromEnv: true
  secretID: MyCloudSecretId
  secretKey: MyCloudSecretKey
policies:
  - namespace: redis_mem_edition    # 云产品类型, https://cloud.tencent.com/document/product/248/50397 查看全部
    conditionID:  8362782           # 触发条件模板 id，https://console.cloud.tencent.com/monitor/alarm/template 查看
    noticeIDs:                      # 通知渠道模板 id，https://console.cloud.tencent.com/monitor/alarm/notice 查看
      - notice-4uks09xe
    name: 基础Redis实例告警         # 告警策略名称，创建时候使用，更新时候也会用这个查询后修改
    remark: 维护 by tcop cmd         # 备注
    tags:                           # 监控策略的 tag 用于管理监控，监控的实例无关
      - key: 用途
        value: 魔法
resourceTags:                       # 监控覆盖的资源标签的 tag
  - key: 用途
    values:
      - 魔法
      - 测试
```