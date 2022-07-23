# Shell集锦

## jq的使用
- jq -r .data.message，可以去掉message外层的双引号

## curl使用

- curl -s 静默模式，可不展示请求进度

## 定时提交

```shell
#!/bin/sh

cd /home/pi/github/blog
git add .
COMMENT=自动同步github[$(date "+%Y-%m-%d %H:%M:%S")]
git commit -m"$COMMENT"
git push origin master
```

## 批量替换多个文件内容
```sh
命令格式：sed -i "s/原内容/新内容/g" `grep 原内容 -rl 所在目录`    注：千万注意这个符号【`】，是【最左上角】那个符号不是单引号
```

## 批量修改文件夹下的文件扩展名

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