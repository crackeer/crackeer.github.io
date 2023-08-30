# VSCode插件

## 1. Golang自动注释插件

> **GoComment**使用`Ctrl + Alt + /`

## 2. AI代码补全

> Tabnine AI Autocomplete for Javascript, Python, Typescript

## 3. Markdown Previvew Enhanced

## 4. Paste Image

# 树莓派

## 设置静态IP
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

# SQLite使用

> SQLite教程：https://www.runoob.com/sqlite/sqlite-tutorial.html

- 删除Table

```sql
DROP TABLE table_name;
drop table if exists table_name;
```

- 修改table名字

```sql
ALTER TABLE UserInfo RENAME TO NewUserInfo;
```

- 增加字段

```sql
ALTER TABLE UserInfo ADD COLUMN Sex Text NOT NULL;
```

- 查看数据表字段

```sql
PRAGMA TABLE_INFO (UserInfo);
```

- 删除表字段

> https://blog.csdn.net/gyymen/article/details/53534267
            
# Lua学习

- lua教程：https://www.runoob.com/lua/lua-tutorial.html
- Lua中时间戳：https://blog.csdn.net/auspark/article/details/103027234


# Docker使用

- docker构建之前需要根据platform分别构建：<a href="/page/code.html?file=/docs/develop/sample/docker_build.sh&title=docker构建脚本" target="_blank">构建的shell脚本</a>

- 一份Dockerfile：<a href="/page/code.html?file=/docs/develop/sample/dockerfile&title=一份dockerfile配置" target="_blank">查看dockerfile文件</a>