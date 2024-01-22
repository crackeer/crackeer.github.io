/// [write_media_file]
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
/// [write_media_file]

/// [simple_read_dir]
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
/// [simple_read_dir]


/// [check-usb]
use std::ffi::CStr;
use libc::{dev_t, ioctl};
use udev::Context;
 
fn main() {
    let context = Context::new().unwrap();
    
    // 获取所有设备节点
    for device in context.enumerate_devices(Some("block")).expect("Failed to enumerate devices.") {
        if let Some(device) = device.into_result() {
            // 判断设备类型为块设备并且是否存在子系统属性
            if device.property_value("ID_BUS").is_some() && device.property_value("SUBSYSTEM").is_some() {
                // 获取设备名称
                let name = CStr::from_ptr((*device.syspath()).as_ref());
                
                println!("Device inserted: {}", name);
                
                // 如果需要进行更多操作，比如读写文件等，可以根据设备路径打开相应的文件或者调用其他函数
            }
        }
    }
}
/// [check-usb]
