## nextjs静态化部署 nginx 配置

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

## nginx设置跨域

```conf
add_header 'Access-Control-Allow-Origin' "$http_origin";
add_header 'Access-Control-Allow-Headers' 'accept,os,accesstoken,content-Type,X-Requested-With,Authorization,apptype,appkey,devid,token,uid,versioncode,versionname,mfg,x-request-id,x-request-uid';
add_header 'Access-Control-Max-Age' '2592000';
add_header 'Access-Control-Allow-Methods' 'GET, PUT, OPTIONS, POST, DELETE';
add_header 'Access-Control-Allow-Credentials' 'true';
add_header 'X-Content-Type-Options' 'nosniff';
add_header 'X-XSS-Protection' '1; mode=block';
add_header 'X-Frame-Options' 'SAMEORIGIN';
```
