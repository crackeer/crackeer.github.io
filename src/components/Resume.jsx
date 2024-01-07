import React from 'react';
import { Tabs, Card, Timeline, QRCode, Tag  } from 'antd';
import Code from './Code';
import WeiboIcon from './WeiboIcon';
import BeikeIcon from './BeikeIcon';
import GanjiIcon from './GanjiIcon';
import HighlightText from "./HighlightText";

const GowayCode = `package main
import (
	"fmt"
	"github.com/crackeer/goaway/container"
	_ "github.com/crackeer/goaway/examples/sign" // 自定义实现签名算法
	"github.com/crackeer/goaway/server"
	"github.com/gin-gonic/gin"
)
func init() {
    // 自定义注册路由
    container.RegisterNakedRouter("access/token", func(ctx *gin.Context) {
		fmt.Println("access/token")
	})
}
func main() {
	server.Main()
}`

const highlightStyle = {
    color: '#bf5726',
    fontSize: '15px',
}

const searchWords = [
    "golang", "redis", "kafka", "贝壳",  "微博", "gin", "gorm", "OpenAPI"
]

const WorkList = [
    {
        children: <Card hoverable bodyStyle={{ padding: '0 10px', width: '80%' }} title={<>
            <BeikeIcon height={30} width={30} />贝壳 · 如视 / <Tag color="#55acee">Golang资深</Tag> / 2020.07 至 今
        </>}>
            <ul>
                <li>负责基础组件APPGateway的设计、开发、</li>
                <li>参与现有项目维护，不断提升服务性能</li>
                <li>参与调研新业务产品需求的技术方案，确立最优方案，并实施落地</li>
            </ul>
            <h4><strong>项目1：</strong>作者 · APPGateway · golang / 动态路由 / 统一验签 / 缓存 / IP频控</h4>
            <ul>
                <li><HighlightText searchWords={searchWords} highlightStyle={highlightStyle}>支撑业务：如视开放平台OpenAPI、内部统一网关</HighlightText></li>
                <li>路由数量：注册路由1000+，注册服务200+，内部API调用</li>
                <li>业务量：日请求量2亿次，Max QPS：3000</li>
            </ul>
            <h4><strong>项目2：</strong>作者 · 任务调度平台 · golang / 任务编排 / 消息</h4>
            <ul>
                <li>支撑业务：未来家AI设计方案产出、VR生产流程、本地化部署基础服务</li>
                <li>任务量：任务类型40+，日处理任务量2w+</li>
                <li><HighlightText searchWords={searchWords} highlightStyle={highlightStyle}>特点1：OpenAPI任务创建、查询、任务统计、Callback、任务重试</HighlightText></li>
                <li>特点2：一定程度的任务编排能力、消息通知订阅</li>
            </ul>
        </Card>
    },
    {
        children: <Card client:load hoverable bodyStyle={{ padding: '0 10px' }} title={<>
            <WeiboIcon height={30} width={30} />微博 · 广告 /<Tag color="#55acee">PHP · Golang工程师 </Tag> / 2018.04 ~ 2020.07
        </>} >
            <ul>
                < li > 介绍：`微博广告业abtest平台，提高广告投放效率</li>
            </ul>
            <h4><strong>技术栈：</strong>php-lavavel，Golang/Gin，GROM</h4>
            <ul>
                <li><HighlightText searchWords={searchWords} highlightStyle={highlightStyle}>后端重构，php迁移Golang，重写接口100+</HighlightText></li>
                <li><HighlightText searchWords={searchWords} highlightStyle={highlightStyle}>实验放量精细化，最大化降低由于实验给广告投放带来的经济损失</HighlightText></li>
            </ul>
        </Card>,
        color: "red"
    },
    {
        children: <Card hoverable bodyStyle={{ padding: '0 10px' }} title={<>
            <GanjiIcon height={35} width={35} /> 赶集 · 斗米招聘 / PHP工程师 / 2015.07~2018.04
        </>}
            style={{ marginTop: '10px' }}>
            <ul>
                <li>介绍：房产无线 / 斗米招聘（独立分拆）</li>
            </ul>
            <h4><strong>技术栈：</strong>php、lavavel，MySQL、Redis</h4>
            <ul>
                <li>维护房产无线端业务正常运行+</li>
                <li>商户招聘管理后台商业化</li>
            </ul>
        </Card>,
        color: 'green'
    }
]

const BaseInfo = () => {
    return <>
        <Card client:idle hoverable bodyStyle={{ padding: '0 10px' }}>
            <ul>
                <li><strong>1993年：</strong> 男 · 湖北十堰</li>
                <li><strong>2011-2015：</strong> 长安大学(211) · 软件工程 · 本科</li>
                <li> <strong>联系方式：</strong>
                    <a href="tel:15611103288">15611103288</a> / <a href="mailto:seekh@qq.com">seekh@qq.com</a>
                </li>
                <li><strong>期望职位：</strong>Golang工程师（武汉）</li>
            </ul>
        </Card>
        <h3>专业技能</h3>
        <Card client:load hoverable bodyStyle={{ padding: '0 10px' }}>
            
            <ul>
                <li>Golang 4年工作经验，有大流量、高并发服务实战经验</li>
                <li>有大型系统的重构经验、以及语言迁移经验</li>
                <li>熟悉Golang标准库，以及常见的性能优化</li>
                <li><HighlightText searchWords={searchWords} highlightStyle={highlightStyle}>熟悉Gin / GORM / Kafka / Bleve等优秀开源库使用</HighlightText></li>
                <li>熟悉MySQL、Redis的使用，并了解相关运行原理和调优机制，实战经验丰富</li>
                <li>熟悉kafka消息队列，有实际业务应用经验</li>
            </ul>
        </Card>
        <h3>个人评价</h3>
        <Card client:load hoverable bodyStyle={{ padding: '0 10px' }}>
            <ul>
                <li>学习能力强，乐于学习新技能，兴趣编程语言：NodeJS / Rust</li>
                <li>沟通能力良好，工作主动、认真，较强的抗压能力</li>
                <li>对代码质量有追求，善于面对需求变化，拥抱变化</li>
            </ul>
        </Card>
    </>
}

const TabLabel = (props) => {
    return <strong style={{fontSize:'16px'}}>{props.children}</strong>
}

const items = [
    {
        key: '1',
        label: <TabLabel>基本信息</TabLabel>,
        children: <BaseInfo />
    },
    {
        key: '2',
        label: <TabLabel>工作经历</TabLabel>,
        children: <>
            <Timeline
                items={WorkList}
                style={{ paddingLeft: '20px' }}
            />
        </>
    },
    {
        key: '3',
        label: <TabLabel>Github项目</TabLabel>,
        children: <>
            <Card hoverable bodyStyle={{ padding: '0 10px' }} title={<>
                app-gateway(Golang) / <a href="https://github.com/crackeer/go-gateway"
                        target="_blank">https://github.com/crackeer/go-gateway</a>
            </>}>
                <strong>main.go文件</strong>
                <Code lang='go'>
                    {GowayCode}
                </Code>
                <ul>
                    <li>【Golang】caddy插件-文件上传：<a href="https://github.com/crackeer/caddy-upload2dir"
                        target="_blank">https://github.com/crackeer/caddy-upload2dir</a> </li>
                    <li>【Golang】caddy插件-数据库读写：<a href="https://github.com/crackeer/caddy-database"
                        target="_blank">https://github.com/crackeer/caddy-database</a> </li>
                    <li>【caddy应用】markdown文档web版：<a href="https://github.com/crackeer/markdown-web"
                        target="_blank">https://github.com/crackeer/markdown-web</a> </li>
                </ul>
            </Card>
        </>
    }
];
const App = () => (
    <Tabs
        defaultActiveKey="1"
        items={items}
    />
);
export default App;