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