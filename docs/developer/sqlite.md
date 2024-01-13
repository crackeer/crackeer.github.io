# SQLite冲啊
> sqlite是我喜欢的数据库啊


## SQLite使用

> SQLite教程：https://www.runoob.com/sqlite/sqlite-tutorial.html

- 删除Table

```sql
DROP TABLE table_name;
drop table if exists table_name;
```

- 修改table名字

```sql
ALTER TABLE UserInfo RENAME TO NewUserInfo;
```

- 增加字段

```sql
ALTER TABLE UserInfo ADD COLUMN Sex Text NOT NULL;
```

- 查看数据表字段

```sql
PRAGMA TABLE_INFO (UserInfo);
```

- 删除表字段

> https://blog.csdn.net/gyymen/article/details/53534267