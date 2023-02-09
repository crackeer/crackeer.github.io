## 刘虎 - 15611103288
- 男 - 1993 - **长安大学本科(211)软件工程**
- 邮箱：767856178@qq.com 


## 工作经历

#### 1. [贝壳找房·如视](https://home.realsee.com/)`/`2020.07 ~ 至今

<!-- tabs:start -->

#### **应用网关 - AppGateway **

**特性：**
- 支持路由类型：**relay**`/`**mesh**`/`**static**三种类型
- **动态路由**`/`**统一验签**`/`**缓存**`/`**IP频控**
- 可视化UI控制后台：在线增加路由、缓存时间调整、频控调整

**技术栈：**

- 后端Golang：
  - [Server路由:gin](https://github.com/gin-gonic/gin) `/` [ORM:gorm](https://github.com/go-gorm/gorm) `/` [JSON解析:tidwall/gjson](https://github.com/tidwall/gjson) 
  - [日志:sirupsen/logrus](https://github.com/sirupsen/logrus) `/` [缓存:go-redis](http://github.com/go-redis/redis/v8) `/` [监控:prometheus](http://github.com/prometheus/client_golang) 
- 前端React：
  - [框架:next.js](https://nextjs.org/docs) `/` [UI:ant-design](https://ant.design/docs/react/introduce-cn) `/` [HTTP请求:axios](https://github.com/axios/axios) 
  - [时间:dayjs](https://day.js.org/en/) `/` [JSON展示:react-json-view](https://github.com/mac-s-g/react-json-view) `/` [JSON编辑:jsoneditor](https://jsoneditoronline.org/)

**成果：**
- 支撑如视开放平台[OpenAPI](https://open-platform.realsee.com/developer/open/api#/)
- 注册路由`1000+`，注册服务`200+`，内部API调用
- 日请求量`2亿次`，MaxQPS:`3000`

**个人实现Github:**

https://github.com/crackeer/go-gateway

#### **任务调度平台 - 牧羊人**

**特性：**
- `开放API`任务创建、查询、任务统计、Callback
- 一定程度的`任务编排`能力、消息通知订阅
- 任务失败率`报警`
- 可视化任务管理后台，实时查看任务执行情况

**架构示意图：**
<img src="/images/shepherd.png" height="310px" width="530px"/>

**技术栈：**

- 后端Golang：
  - [Server路由:gin](https://github.com/gin-gonic/gin) `/` [ORM:gorm](https://github.com/go-gorm/gorm) `/` [消息队列:kafka-go](https://github.com/segmentio/kafka-go) 
  - [日志:sirupsen/logrus](https://github.com/sirupsen/logrus) `/` [缓存:go-redis](http://github.com/go-redis/redis/v8) 
- 前端React：
  - [框架:next.js](https://nextjs.org/docs) `/` [UI:ant-design](https://ant.design/docs/react/introduce-cn) `/` [HTTP请求:axios](https://github.com/axios/axios) 
  - [时间:dayjs](https://day.js.org/en/) `/` [JSON展示:react-json-view](https://github.com/mac-s-g/react-json-view) `/` [JSON编辑:jsoneditor](https://jsoneditoronline.org/)

**成果：**
- 支撑业务：未来家`AI设计`方案产出、`VR生产流程`、`本地化部署`基础服务
- 任务数量40+，日处理任务量`2w+`

<!-- tabs:end -->

#### 2. 微博·广告 `/` 2018.04 ~ 2020.07

<!-- tabs:start -->

#### **abtest平台**

**介绍：**
微博广告业务部的abtest平台，服务于PM、RD，致力于提高算法策略迭代效率

**技术栈：**
`php-lavavel`，`Golang/Gin`，`GROM`、Clickhouse、`Vue`、`Element-UI`

**成果：**
- 从0重构后端架构，由php转移到go，使用gin框架，累计重写接口100+
- 参与实验数据存储方案设计，并提供实验多版本、多维度查询接口，实验效果接口耗时缩短到20s内
- 实验指标管理UI化，降低PM录入指标门槛
- 实验放量精细化管控，最大化降低由于实验给广告投放带来的经济损失

<!-- tabs:end -->

## 技能

<!-- tabs:start -->

#### **Golang**

- 使用golang语言，并能编写稳定、高并发、安全的代码
- gin / gorm /

<!-- tabs:end -->
