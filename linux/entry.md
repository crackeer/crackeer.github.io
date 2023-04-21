## shell中解析json: jq
- jq -r .data.message，可以去掉message外层的双引号

## curl静默模式

- curl -s 静默模式，可不展示请求进度

## git定时提交

```shell
#!/bin/sh

cd /home/pi/github/blog
git add .
COMMENT=自动同步github[$(date "+%Y-%m-%d %H:%M:%S")]
git commit -m"$COMMENT"
git push origin master
```

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
```sh
#!/bin/bash
function red() {
    echo -e "\033[31m$1\033[0m"
}

function yellow() {
    echo -e "\033[33m$1\033[0m"
}

function green() {
    echo -e "\033[32m$1\033[0m"
}

CURRENT_BRANCH=$(git symbolic-ref -q --short HEAD);
if [[ $CURRENT_BRANCH =~ "test" ]]; then 
    red "无法将`test`合并到其他分支"
    exit 1
fi

if [[ $CURRENT_BRANCH =~ "preline" ]]; then 
    red "无法将`preline`合并到其他分支"
    exit 1
fi

green "Current branch: $CURRENT_BRANCH"

if [ -n "$1" ]; then
    MERGE_TO_BRANCH=$1
else
    red "Please specify merge to branch"
    exit 1
fi

if [ "$CURRENT_BRANCH"x = "$MERGE_TO_BRANCH"x ]; then
    red "can't  merge to self"
    exit 1
fi

CURRENT_BRANCH=$(git symbolic-ref -q --short HEAD);
green "Merge: $CURRENT_BRANCH -> $MERGE_TO_BRANCH";
echo ""

yellow "1、切换$MERGE_TO_BRANCH"
git checkout $MERGE_TO_BRANCH
echo ""

yellow "2、Pull $MERGE_TO_BRANCH 代码"
git pull origin $MERGE_TO_BRANCH
echo ""

yellow "3、Pull $CURRENT_BRANCH 代码"
git pull origin $CURRENT_BRANCH
echo ""

yellow "4、Push To $MERGE_TO_BRANCH"
git push origin $MERGE_TO_BRANCH
echo ""

yellow "切换回 $CURRENT_BRANCH"
git checkout $CURRENT_BRANCH
echo ""

green "Successfully"
echo ""

```

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


