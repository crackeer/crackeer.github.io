# 准备脚本

```shell
#!/bin/bash

# 配置变量
LOCAL_HOST="127.0.0.1"
LOCAL_USER="主机用户"
LOCAL_PASSWORD="主机密码"
LOCAL_DB_NAME="数据库"
REMOTE_USER="目标主机用户"
REMOTE_HOST="目标主机IP"
REMOTE_PASSWORD="目标主机密码"
REMOTE_PATH="目标主机备份地址"

# 创建备份文件名
TIMESTAMP=$(date +"%F")
BACKUP_FILE="db_backup_${TIMESTAMP}.sql"

# 导出数据库到本地
mysqldump -u ${LOCAL_USER} -p${LOCAL_PASSWORD} ${LOCAL_DB_NAME} --ignore-table=需要忽略的表（可以不写） --ignore-table=多个表时就这样写 > ${BACKUP_FILE}
```

# 增加执行权限

```shell
chmod -R 775 backup.sh
```

