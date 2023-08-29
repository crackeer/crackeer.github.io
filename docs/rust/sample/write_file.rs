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