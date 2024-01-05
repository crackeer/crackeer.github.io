import React from 'react';
import { Tabs, Card, Timeline, QRCode } from 'antd';
const onChange = (key) => {
    console.log(key);
};

const WorkList = [
    {
        children: <Card hoverable bodyStyle={{ padding: 0, width: '80%' }} title="2020.07 至 今">
            <ul>
                <li>负责基础组件APPGateway的设计、开发、</li>
                <li>参与现有项目维护，不断提升服务性能</li>
                <li>参与调研新业务产品需求的技术方案，确立最优方案，并实施落地</li>
            </ul>
            <h4><strong>项目1：</strong>作者 · APPGateway · golang / 动态路由 / 统一验签 / 缓存 / IP频控</h4>
            <ul>
                <li>支撑业务：如视开放平台OpenAPI、内部统一网关</li>
                <li>路由数量：注册路由1000+，注册服务200+，内部API调用</li>
                <li>业务量：日请求量2亿次，Max QPS：3000</li>
            </ul>
            <h4><strong>项目2：</strong>作者 · 任务调度平台 · golang / 任务编排 / 消息</h4>
            <ul>
                <li>支撑业务：未来家AI设计方案产出、VR生产流程、本地化部署基础服务</li>
                <li>任务量：任务类型40+，日处理任务量2w+</li>
                <li>特点1：开放API任务创建、查询、任务统计、Callback、任务重试</li>
                <li>特点2：一定程度的任务编排能力、消息通知订阅</li>
            </ul>
        </Card>,
        color: "green"
    },
    {
        children: <Card client:load hoverable bodyStyle={{ padding: 0 }} title="微博·广告 / PHP·Golang工程师 / 2018.04 ~ 2020.07">
            <ul>
                <li>介绍：`微博广告业abtest平台，提高广告投放效率</li>
            </ul>
            <h4><strong>技术栈：</strong>php-lavavel，Golang/Gin，GROM</h4>
            <ul>
                <li>后端重构，php迁移Golang，重写接口100+</li>
                <li>实验放量精细化，最大化降低由于实验给广告投放带来的经济损失</li>
            </ul>
        </Card>
    },
    {
        children: <Card hoverable bodyStyle={{ padding: 0 }} title=" 赶集·斗米招聘 / PHP工程师 / 2015.07~2018.04"
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
        color: 'gray'
    }
]
const items = [
    {
        key: '1',
        label: '基本信息',
        children: <>
            <Card client:idle hoverable bodyStyle={{ padding: 0 }}>
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
            <Card client:load hoverable bodyStyle={{ padding: 0 }}>
                <ul>
                    <li>Golang 4年工作经验，有大流量、高并发服务实战经验</li>
                    <li>有大型系统的重构经验、以及语言迁移经验</li>
                    <li>熟悉Golang标准库，以及常见的性能优化</li>
                    <li>熟悉Gin / GORM / Kafka / Bleve等优秀开源库使用</li>
                    <li>熟悉MySQL、Redis的使用，并了解相关运行原理和调优机制，实战经验丰富</li>
                    <li>熟悉kafka消息队列，有实际业务应用经验</li>
                </ul>
            </Card>
            <h3>个人评价</h3>
            <Card client:load hoverable bodyStyle={{ padding: 0 }}>
                <ul>
                    <li>学习能力强，乐于学习新技能，兴趣编程语言：NodeJS / Rust</li>
                    <li>良好的沟通能力，工作积极主动、认真负责，并有较强的抗压能力</li>
                    <li>对代码质量有追求，善于面对需求变化，拥抱变化</li>
                </ul>
            </Card>
        </>
    },
    {
        key: '2',
        label: '工作经历',
        children: <>
            <Timeline
                items={WorkList}
                style={{ paddingLeft: '20px' }}
            />
        </>
    },
    {
        key: '3',
        label: 'Github项目',
        children: <>
            <Card hoverable bodyStyle={{ padding: 0 }} title="Go版本的网关">
                <code>
                    sakjgsjkg
                </code>
                <ul>
                    <li>【Golang】Gateway：<a href="https://github.com/crackeer/go-gateway"
                        target="_blank">https://github.com/crackeer/go-gateway</a> </li>
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
        onChange={onChange}
    />
);
export default App;