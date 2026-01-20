# Linux终端环境下进行802-1x网络认证

# 背景

> 有台工作站，安装了UOS，插上公司内网网线，无法上网，需要802.1x认证



## 添加有线

```shell
sudo nmcli con add type ethernet \
    con-name "有线-8021X" \
    ifname eno1 \
    ipv4.method auto \
    ipv6.method auto \
    802-1x.eap peap \
    802-1x.phase2-auth mschapv2 \
    802-1x.identity "liuhu016" \
    802-1x.password "XXXXXXXX"
    
 sudo nmcli con up "有线-8021X"
```



## 修改密码

- 展示

```shell
nmcli con show
```

- 修改密码

```shell
sudo nmcli con modify "有线-8021X" 802-1x.password "your_new_password"
```

