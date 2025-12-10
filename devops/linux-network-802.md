# Linux终端环境下进行802-1x网络认证

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

