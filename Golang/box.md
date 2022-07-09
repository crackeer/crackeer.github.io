# 工具箱

## awsome-go：https://awesome-go.com/

## 很好用的模板引擎

github地址：https://http://github.com/flosch/pongo2/v4

使用示例：

```go
tpl, err := pongo2.FromString(content)
if err != nil {
    panic(err)
}
data := map[string]interface{}
out, err := tpl.Execute(pongo2.Context(data))
if err != nil {
    mailer.Failure(err.Error())
    return
}
```

## 验证工具

- https://github.com/gookit/validate?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego




