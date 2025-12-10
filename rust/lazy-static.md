# 使用lazy-static修改获取全局变量

## Cargo.toml

```toml
[dependencies]
lazy_static = "1.4.0"
```

## lib.rs

```rust
#[macro_use]
extern crate lazy_static;
```

## logic.rs

```rust
use std::sync::{Arc, Mutex};
lazy_static! {
   pub static ref SESSION_MAP: Arc<Mutex<HashMap<String, Session>>> = Arc::new(Mutex::new(HashMap::new()));
}
```
