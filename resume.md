## 刘虎 - 15611103288
- 男 - 1993 - 767856178@qq.com 
- 教育背景：长安大学(211) - 软件工程 - 2011至2015
- 求职意向：武汉 - Golang开发

## 一、专业技能


- Golang 4年工作经验，有大流量、高并发服务实战经验
- 有大型系统的重构经验、以及语言迁移经验
- 熟悉Golang标准库，以及常见的性能优化
- 熟悉Gin / GORM / Kafka / Bleve等优秀开源库使用
- 熟悉MySQL、Redis的使用，并了解相关运行原理和调优机制，实战经验丰富
- 熟悉kafka消息队列，有实际业务应用经验


## 二、工作经历

### 1. [贝壳找房·如视](https://home.realsee.com/)`/`2020.07 ~ 至今 `/` Golang工程师

###### 工作职责：

- 负责基础组件`APPGateway`的设计、开发、
- 参与现有项目维护，并不断提升服务性能
- 参与调研新业务产品需求的技术方案，确立最优方案，并实施落地

###### 项目经验：

<!-- tabs:start -->

#### **应用网关 - APPGateway**

**本人角色**

`主要开发者`

**特性：**
- 支持路由类型：**relay**`/`**mesh**`/`**static**三种类型
- **动态路由**`/`**统一验签**`/`**缓存**`/`**IP频控**
- 可视化UI控制后台：灵活控制路由、**服务超时**等

**技术栈：**

- 后端Golang：[Server路由:gin-gonic/gin](https://github.com/gin-gonic/gin)`/`[ORM:go-gorm/gorm](https://github.com/go-gorm/gorm)`/`[JSON解析:tidwall/gjson](https://github.com/tidwall/gjson)`/`[缓存:go-redis](http://github.com/go-redis/redis/v8) `/` [监控:prometheus](http://github.com/prometheus/client_golang) 

**成果：**
- 支撑如视开放平台[OpenAPI](https://open-platform.realsee.com/developer/open/api#/)
- 注册路由`1000+`，注册服务`200+`，内部API调用
- 日请求量`2亿次`，MaxQPS:`3000`

**个人实现Github:**

https://github.com/crackeer/go-gateway

#### **任务调度平台 - 牧羊人**


**本人角色**

`主要开发者`

**特性：**

- 系统特点1：开放API任务创建、查询、任务统计、Callback、任务重试
- 系统特点2：一定程度的任务编排能力、消息通知订阅

**架构示意图：**
<img src="/images/shepherd.png" height="310px" width="530px"/>

**技术栈：**

- 后端Golang：[Server路由:gin-gonic/gin](https://github.com/gin-gonic/gin) `/` [ORM:go-gorm/gorm](https://github.com/go-gorm/gorm) `/` [消息队列:segmentio/kafka-go](https://github.com/segmentio/kafka-go) 

**成果：**
- 支撑业务：未来家`AI设计`方案产出、`VR生产流程`、`本地化部署`基础服务
- 任务数量40+，日处理任务量`2w+`

<!-- tabs:end -->

### 2. 微博·广告 `/` 2018.04 ~ 2020.07 `/` PHP-Golang工程师

<!-- tabs:start -->

#### **ABTest平台**

**介绍：**`微博广告业务的abtest平台，致力于提高广告投放效率

**技术栈：**`php-lavavel`，`Golang/Gin`，`GROM`，`Vue/Element-UI`

**成果：**
- 后端重构，php迁移Golang，重写接口100+
- 成为衡量广告策略的重要平台
- 实验放量精细化管控，最大化降低由于实验给广告投放带来的经济损失

<!-- tabs:end -->

## 三、个人评价

- 学习能力强，乐于学习新技能，在学Rust语言(目标：用Rust重写go版的APPGateway)
- 有良好的沟通能力，工作积极主动、认真负责，并有较强的抗压能力
- 对代码质量有追求，善于面对需求变化，拥抱变化
- Show Me Your Code
  - 网关gateway：https://github.com/crackeer/go-gateway
  - 制作Web页：https://github.com/crackeer/goweb