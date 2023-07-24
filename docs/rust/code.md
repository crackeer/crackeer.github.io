## 创建文件夹 & 写文件

```rust
pub fn write_media_file(dir: String, name : String, content: Vec<u8>) -> String {
    let mut tmp_path = std::path::Path::new(&dir);
    if let Err(err) = std::fs::create_dir_all(&tmp_path) {
        return err.to_string();
    }
    let path_buf = tmp_path.join(&name);
    std::fs::write(tmp_path, &content.as_slice());
    if let Ok(mut file) = File::create(path_buf.as_path()) {
        if let Ok(_) = file.write_all(&content) {
            return String::from("ok")
        }
    }
    String::from("error")
}
```

## 读取文件夹

```rust
pub fn simple_read_dir(dir: String, ext: String) -> Vec<FileItem> {
    let mut list: Vec<FileItem> = Vec::new();
    let entry = read_dir(dir);
    if entry.is_err() {
        return list;
    }
    for item in entry.unwrap().into_iter() {
        if let Err(_) = item {
            continue;
        }
        let file_name = item.as_ref().unwrap().file_name().into_string().unwrap();
        let entry = item.unwrap().metadata().unwrap();
        if entry.is_dir() {
            if !file_name.starts_with(&".") {
                list.push(FileItem {
                    path: file_name,
                    item_type: String::from("dir"),
                })
            }
        } else {
            if file_name.ends_with(&ext) {
                list.push(FileItem {
                    path: file_name,
                    item_type: String::from("file"),
                })
            }
        }
    }
    list
}
```

## 使用ssh2链接服务器
- convert ssh id_rsa to PEM Format
https://github.com/alexcrichton/ssh2-rs
```
ssh-keygen -p -m PEM -f ~/.ssh/id_rsa
```    

