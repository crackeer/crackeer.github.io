# Let's Rust

## 一些个人项目

- tauri写的一个简单的markdown编辑器：https://github.com/crackeer/tauri-markdown


## 使用ssh2登录服务器
> https://github.com/alexcrichton/ssh2-rs

- convert ssh id_rsa to PEM Format

```sh
ssh-keygen -p -m PEM -f ~/.ssh/id_rsa
```    

## 调用C静态链接库
- https://blog.csdn.net/ytxwhwlb/article/details/103465066

## 检测USB

[check-usb](./code/rust.rs ':include :type=code :fragment=check-usb')


## 读取文件夹

[simple_read_dir](./code/rust.rs ':include :type=code :fragment=simple_read_dir')

## Tauri使用问题

### release版本可以检查元素

```
tauri = { version = "1.0.0-rc.8", features = ["api-all", "devtools"] }
```

## 更新rust版本:`rustup update`,速度很慢

需要设置proxy
- RUSTUP_DIST_SERVER=https://rsproxy.cn
- RUSTUP_UPDATE_ROOT=https://rsproxy.cn/rustup

