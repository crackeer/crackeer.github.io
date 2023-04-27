# nextjs静态化部署nginx配置
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