# 一段js代码遍历

> 抽取中文

```js
var jsonData = {
    "abc" : "simple",
    "json" : [
        {
            "name" : "加氨氮"
        },
        {
            "name" : "简单"
        }
    ],
    "Ok" : false,
    "age" : 99,
    "name" : "中文字符的正"
}

function containsChinese(str) {  
    // 匹配中文字符的正则表达式  
    var reg = /[\u4e00-\u9fa5]/gm;  
    return reg.test(str);  
}


var extractChineseFields = (data, prefix) => {
    let retData = {}
    if (typeof data == 'object' && data.length == undefined) {
        Object.keys(data).forEach(key => {
            let tmpKey = prefix.length > 0 ? prefix+'.'+key : key
            if(typeof data[key] == 'string' && containsChinese(data[key])) {
                retData[tmpKey] = data[key]
            } else if (typeof data[key] == 'object') {
                let tmpData = extractChineseFields(data[key], tmpKey)
                Object.keys(tmpData).forEach(kk => {
                    retData[kk] = tmpData[kk]
                })
            }
        })
    }
    if (typeof data == 'object' && data.length != undefined) {
        for(var i in data) {
            let tmpKey = prefix.length > 0 ? prefix+'.'+i : key
            let tmpData = extractChineseFields(data[i], tmpKey)
            Object.keys(tmpData).forEach(kk => {
                retData[kk] = tmpData[kk]
            })
        }
    }
    return retData
}

let retData = extractChineseFields(jsonData, "")
console.log(retData)
```

