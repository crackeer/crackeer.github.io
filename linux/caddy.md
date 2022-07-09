# Caddy配置

## 一份配置文件

```json
{
	"admin": {
		"disabled": false,
		"listen": "0.0.0.0:2019",
		"enforce_origin": false,
		"origins": [""]
	},
	 "apps":{
        "http":{
            "servers":{
                "static":{
                    "idle_timeout":30000000000,
                    "listen":[
                        "0.0.0.0:2022"
                    ],
                    "max_header_bytes":10240000,
                    "read_header_timeout":10000000000,
                    "routes":[
                        {
                            "handle":[
                                {
                                    "handler":"file_server",
                                    "root" : "/home/pi/",
                                    "browse": {
                                        "template_file": ""
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