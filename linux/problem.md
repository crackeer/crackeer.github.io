## 树莓派设置静态IP
- https://www.cnblogs.com/lfri/p/12115211.html

## 没有屏幕设置wifi

```conf
country=CN
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
 
network={
    ssid="HUAWEI30"
    psk="15611103288"
    priority=3
}
 
network={
    ssid="yuanbao0616"
    psk="zjyx2011"
    priority=4
}
```

## 开启gpio，风扇自动控制

```shell
#!/bin/sh

# export
if [ ! -d "/sys/class/gpio/gpio23/" ];then
    echo 18 > /sys/class/gpio/export
    echo out /sys/class/gpio/gpio23/direction
fi

temp=$(cat /sys/class/thermal/thermal_zone0/temp)
echo $temp 
if [ $temp -gt 39000 ] ; then
    echo 0 > /sys/class/gpio/gpio23/value 
else
    echo 1 > /sys/class/gpio/gpio23/value 
fi

```