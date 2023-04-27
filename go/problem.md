## 私有仓库不支持https

> https://www.cnblogs.com/hiwz/p/12652153.html

## Gorm使用sqlite不开启CGO

> https://github.com/go-gorm/sqlite

```go
import (
  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
)

db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
```