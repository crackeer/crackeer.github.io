
#### http状态码502与504的区别？
1. 502是网关错误，504是网关超时
2. 二者很类似，502是代理服务器后面的真实服务器节点配置出了问题或者已经挂掉了，而504是代理服务器后面的真实服务器已经过载，它要处理的请求报文实在太多，忙不过来了。
3. 502还有一种情况就是nginx与fastcgi即PHP进程配合的不恰当，导致返回502网关错误