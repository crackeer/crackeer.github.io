# 使用Openresty统计pv / uv

## nginx配置如下

```
 server {
    listen 9003;
    server_name _;

    location / {
        log_by_lua_file /etc/openresty/scripts/jinxing-stat.lua;
        ...
    }
}
```

## lua脚本如下

```lua
local function push_to_redis(premature, user_token)
    if premature then return end
    local redis = require "resty.redis"
    local red = redis:new()
    red:set_timeout(1000) 
    local ok, err = red:connect("rfrm-business-redis.redis.svc.cluster.local", 6379)
    if not ok then
        ngx.log(ngx.ERR, "failed to connect to redis: ", err)
        return
    end

    local today = os.date("%Y%m%d")
    red:init_pipeline()
    red:incr("jingxin-stats:pv:" .. today )
    red:pfadd("jingxin-stats:uv:" .. today , user_token)
    red:commit_pipeline()
    red:set_keepalive(10000, 100)
end

local path_whitelist = { 
    ["/home"] = true ,
    ["/collection-3d"] = true
}
local uri = ngx.var.uri
if path_whitelist[uri] then
    local user_token = ngx.req.get_headers()["X-Realsee-Token"] or "anonymous"
    local ok, err = ngx.timer.at(0, push_to_redis, user_token)
    if not ok then
        ngx.log(ngx.ERR, "failed to create timer: ", err)
    end
end
```