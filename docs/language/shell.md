# 常见shell命令使用

## shell中解析json: jq
- jq -r .data.message，可以去掉message外层的双引号

## curl静默模式

- curl -s 静默模式，可不展示请求进度


## sed批量替换
```sh
命令格式：sed -i "s/原内容/新内容/g" `grep 原内容 -rl 所在目录`    注：千万注意这个符号【`】，是【最左上角】那个符号不是单引号
```

## 批量改文件扩展名

```sh
#!/bin/bash
list=$(find . -name '*.html')
for item in $list
    do
        if [ "${item##*/}" == "index.html" ];then
            echo $item;
        else
            mv $item ${item%.*}
        fi
        #
    done
```

## Git合并到其他分支

## 添加system service

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

## clashX设置白名单

打开终端，在新建一个文件
```sh
vim ~/.config/clash/proxyIgnoreList.plist
```
文件内容可以从官方下载。不方便下载的也可以直接复制我下面的文件内容，改一下就好

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<array>
    <string>192.168.0.0/16</string>
    <string>10.0.0.0/8</string>

    <!--上面的不要删，在下面添加你想要忽略的域名-->
    <string>*.lc.com</string>
    <string>*-lc.com</string>
    <string>*-local.com</string>
    <string>*-.local.com</string>
</array>
</plist>
```

## shell中字符串分割

> 参照：https://blog.csdn.net/Baiyi_destroyer/article/details/126530139
- 1. ${parameter//pattern/string},用string来替换parameter变量中所有匹配的pattern

```sh
#!/bin/bash
string="hello,shell,split,test"  
array=(${string//,/ }) 
for var in ${array[@]}
do
   echo $var
done 

```

- 自定义IFS变量, 改变分隔符, 对字符串进行切

```sh
#!/bin/bash
string="hello,shell,split,test"  
#对IFS变量 进行替换处理
OLD_IFS="$IFS"
IFS=","
array=($string)
IFS="$OLD_IFS"
for var in ${array[@]}
do
   echo $var
done

```

- 使用tr：由于只是对单个字符进行的替换，则可以用  echo args |   tr "oldSpilt" "newSpilt"  的方式实现。

```sh
#!/bin/bash
 
string="hello,shell,split,test"  
array=(`echo $string | tr ',' ' '` )  
 
for var in ${array[@]}
do
   echo $var
done 

```

## 字符串切割成数组

> https://www.cnblogs.com/linux985/p/14866985.html

```sh
OLD_IFS="$IFS" 
IFS="," 
arr=($a) 
IFS="$OLD_IFS" 
for s in ${arr[@]} 
do 
    echo "$s" 
done
```

## ssh查看连接过程

```sh
 ssh -vT git@github.com
 ```


