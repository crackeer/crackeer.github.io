# QA问答

            
## Lua学一下

- lua教程：https://www.runoob.com/lua/lua-tutorial.html
- Lua中时间戳：https://blog.csdn.net/auspark/article/details/103027234


## Caddy配置使用

> 文档：https://caddyserver.com/docs/
- 一份文档

[caddy.json](./code/caddy.json ':include :type=code json')

## Nginx配置

### nextjs静态化部署 nginx 配置

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

## nginx设置跨域

```conf
add_header 'Access-Control-Allow-Origin' "$http_origin";
add_header 'Access-Control-Allow-Headers' 'accept,os,accesstoken,content-Type,X-Requested-With,Authorization,apptype,appkey,devid,token,uid,versioncode,versionname,mfg,x-request-id,x-request-uid';
add_header 'Access-Control-Max-Age' '2592000';
add_header 'Access-Control-Allow-Methods' 'GET, PUT, OPTIONS, POST, DELETE';
add_header 'Access-Control-Allow-Credentials' 'true';
add_header 'X-Content-Type-Options' 'nosniff';
add_header 'X-XSS-Protection' '1; mode=block';
add_header 'X-Frame-Options' 'SAMEORIGIN';
```


## 删除git中大文件+记录
- https://blog.csdn.net/HappyRocking/article/details/89313501

```bash
git rev-list --all | xargs -rL1 git ls-tree -r --long | sort -uk3 | sort -rnk4 | head -10
git filter-branch --tree-filter "rm -f {filepath}" -- --all
git push -f --all
```

## 进程`VS`线程

**进程：** 进程是系统分配资源和调度的基本单位，也就是说进程可以单独运行一段程序。
**线程：** 线程是cpu调度和分派的最小基本单位。

**区别：**

1. 一个进程可以包含至少一个线程，一般来说也就是主线程，而一个线程只能属于一个进程；
2. 进程拥有独立的内存，而线程没有独立的资源空间， 只是暂时存储在计数器，寄存器，栈中，同一个进程间的线程可以共享资源。
3. 将代码放入到代码区之后，进程产生，但还没执行，我们所说的执行一般是是主线程main函数开始执行。
4. 进程比线程更加消耗资源
5. 进程对资源的保护要求高，而线程要求不高
6. 进程是处理器这一层面的抽象，而线程是进程的基础上进一步并发的抽象
7. 同一个进程下，一个线程的挂掉，会导致整个进程的挂掉，而进程之间不会相互影响
8. 总的来说：我们都知道程序不能单独运行，只有将它放入内存中，分配资源才能运行，程序是指令的集合，而进程是程序的一次执行活动，属于动态概念
9. 我们可以打个比方:进程相当于某一个大型项目，世界上可能有人同时在做这个项目，有其独特的方式；而线程就相当于这个项目下的一些程序员，多个程序员去完成这一个项目肯定要比一个人完成快的多，也就是能在同一时间操作。


## clashX设置白名单

- 打开终端，在新建一个文件

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

## zip分包 / 合并

```sh
# 分
zip -s 100m myfolder.zip myfolder
# 合
zip -s 0 myfolder.zip --out unsplit.zip
```

