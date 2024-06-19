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

