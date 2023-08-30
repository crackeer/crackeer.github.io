# A. Caddy

> 文档：https://caddyserver.com/docs/

- 一份配置：<a href="/page/code.html?file=/docs/sidecar/sample/caddy.json&title=一份caddy.json" target="_blank">查看JSON</a>

----

# B. Nginx

- nextjs 静态化部署 nginx 配置

```nginx
 server {
        listen       8080;
        server_name  localhost;

        #charset koi8-r;
        root /root/your/path/out;
        location ~/page/(.*)$ {
            rewrite ^\/page\/(.*)$ /$1.html break;
        }
        location ~/(.*)$ {
            if (!-e $request_filename) {
                rewrite ^(.*)$ /$1.html break;
            }
            #alias /root/your/path/out/;
        }

    }
```
