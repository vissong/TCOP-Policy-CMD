# TCOP-Policy-CMD
用于管理腾讯云可观测平台告警策略的命令行工具 / a command tool to manage Tencent Cloud Observable Platform alarm policy


ROADMAP：
- 支持拉取告警策略列表 ✅
- 支持搜索告警列表 ✅
- 支持创建告警策略，支持按照标签绑定，支持使用触发模板来设置出发条件 ✅ 
- 支持编辑告警策略，更换触发条件，告警渠道，更改绑定对象的标签 ✅
- 支持基于 yaml 的模板描述，统一刷新告警策略的范围
- 支持触发条件模板的拉取/搜索

不足：
- 目前云接口不支持使用【全部对象】的方式进行资源与告警策略的绑定
- 目前部分云产品的告警策略（比如 Ckafka 的 CousumerGroup+Topic 监控）
不支持标签绑定资源之后配置指标监控，这种建议手动到控制台设置。

config demo：
```yaml
# 监控策略列表
policies:
  - namespace: redis_mem_edition # 云产品类型
    conditionID:  8362782 # 触发条件模板 id
    noticeIDs:
      - notice-4uks09xe # 通知渠道模板 id
    name: QQ基础 Redis 实例告警 # 告警策略名称，创建时候使用，更新时候也会用这个查询后修改
    remark: 维护 by tcop cmd  # 备注
    tags: # 监控策略的 tag
      - key: 用途
        value: 魔法
# 资源标签
resourceTags: 
  - key: 用途
    values:
      - 魔法
      - 测试
```