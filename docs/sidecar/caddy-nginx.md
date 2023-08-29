# Caddy

> https://caddyserver.com/docs/

## 一份Sample配置

```json
{
    "admin": {
        "disabled": false,
        "listen": "0.0.0.0:2019",
        "enforce_origin": false,
        "origins": [
            ""
        ]
    },
   
    "apps": {
        "http": {
            "servers": {
                "static": {
                    "idle_timeout": 30000000000,
                    "listen": [
                        "0.0.0.0:80"
                    ],
                    "max_header_bytes": 10240000,
                    "read_header_timeout": 10000000000,
                    "routes": [
                        {
                            "match": [
                                {
                                   "host":["simple.com"]
                                }
                            ],
                            "handle": [
                                {
                                    "handler": "reverse_proxy",
                                    "upstreams": [
                                        {
                                            "dial": "localhost:9393"
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "match": [{
                                "path_regexp" :
                                {
                                   "name" : "iss",
                                   "pattern" : "^/iss-vr-index"
                                }

                            }],
                            "handle":[
                                {
                                    "handler":"reverse_proxy",
                                    "upstreams" : [
                                        {
                                            "dial" : "10.26.52.27:11111"
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "match": [
                                {
                                   "host":["www.simple.com"]
                                }
                            ],
                            "handle": [
                                {
                                    "handler": "file_server",
                                    "root": "/your/patc",
                                    "browse": {
                                        "template_file": ""
                                    }
                                }
                            ],
                            "terminal" : true
                        },
                        {
                            "match": [
                                {
                                    "header": {
                                        "proxy": [
                                            "simple"
                                        ]
                                    }
                                }
                            ],
                            "handle": [
                                {
                                    "handler": "reverse_proxy",
                                    "dynamic_upstreams": {
                                        "source": "a",
                                        "name": "proxy.host.com"
                                    },
                                    "headers": {
                                        "request" : {
                                            "add" : {
                                                "Host" : ["proxy.host.com"]
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    ]
                }
            }
        }
    }
}

```

# Nginx

## nextjs静态化部署nginx配置

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