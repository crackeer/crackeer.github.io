# 添加system service

```sh
[Unit]
Description=XXXX
Documentation=http://nginx.org/en/docs/
 
[Service]
Type=simple
PIDFile=/run/xxx.pid
ExecStart=/root/.xxx.sh start
ExecReload=/bin/kill -s HUP $MAINPID
ExecStop=/bin/kill -s QUIT $MAINPID
PrivateTmp=true
 
[Install]
WantedBy=multi-user.target
```

