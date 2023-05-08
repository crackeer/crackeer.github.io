## 使用`ali-rds`操作数据

```js
const rds = require('ali-rds');
const onlineDB = rds({
    host: 'what kind of host',
    port: '3306',
    user: 'username',
    password: 'password',
    database: 'database', 
})
var select = async function() {
   let data = await onlineDB.select("table_name", {
       where : {
        "task_id" : '37675849834616860',
       },
       columns: ['author', 'title'],
       orders: [['id', 'desc']],
       limit : 3,
    })
    console.log(data)
}

var execSQL = async function() {
    let res = await db.query("UPDATE `table_name` SET `picture_url` = 'simple_url' WHERE `id` = 1060")
}
// update
var updateData = async function() {
    await onlineDB.update("table_name", {
        config : "Some Config"
        status : 5,
        name : "your name",
    }, {
        where : {
            id : someID
        }
    })
}
// Delete
db.delete('table_name', {
    id : 1022
})
```

## 使用axios请求

```js
var axios = require('axios')
var getData = async function(value) {
    let res = await axios.get('http://i.svc.open.realsee.com/inner/work/show', {
        params : {
            key : value
        },
    })
    return res.data.data
}
getData('54djM0oR6ky3Ey8e')
```

## 读取文件

```js
var fs = require("fs")
module.exports = {
	readFileAsJSON : function(fileName) {
	    let data = fs.readFileSync(fileName)
		return JSON.parse(data.toString())
	},
	readFileAsString : function(fileName) {
	    let data = fs.readFileSync(fileName)
	    return data.toString()
	}
}
```

## MD5加密

```js
var crypto = require("crypto");
module.exports = function md5(s) {
    //注意参数需要为string类型，否则会报错
     return crypto.createHash('md5').update(String(s)).digest('hex');
}
```

## 读取excel

```js
var xlsx = require('node-xlsx');
module.exports = {
	readFileColumn : function(fileName, columnIndex) {
	    var sheets = xlsx.parse(fileName);//获取到所有sheets
	    let obj = {}
	    let list = []
	    for(var rowId in sheets[0]['data']){
	        if(rowId > 0) {
	            var value =sheets[0]['data'][rowId][columnIndex];
	            if(obj[value] == undefined) {
	                obj[value] = true
	                list.push(value)
	            }
	        }
	        
	    }
	    return list
	},
	readFile : function(fileName) {
	    var sheets = xlsx.parse(fileName);//获取到所有sheets
	    let obj = {}
	    let list = []
	    for(var rowId in sheets[0]['data']){
	        if(rowId > 0) {
	            var value =sheets[0]['data'][rowId]
	            if(value.length > 0) {
	            	list.push(value)
	            }
	        }
	    }
	    return list
	}
}
```

## 写excel

```js
const xlsx = require('node-xlsx')

var handle = async () => {
    let xlsxData = []
    xlsxData.push({
        "name" : 'sheet1',
        "data" : [
            ["abbb", "sss", "sss"]
        ]
    })
    xlsxData.push({
        "name" : 'sheet2',
        "data" : [
            ["abbb", "sss", "sss"]
        ]
    })
    const buffer = xlsx.build(xlsxData)
    fs.writeFile('output.xlsx', buffer, (err) => {
        console.log(err)
    })
}
```

## puppeteer爬取页面数据

```js
const puppeteer = require('puppeteer');
(async () => {
    const browser = await puppeteer.launch({ headless: true });
    const page = await browser.newPage();
    let res1 = await page.goto('https://baidu.cn/gXQQ4RdB', {
        waitUntil: 'networkidle2',
    });
    console.log("Finished");
    const aHandle = await page.evaluateHandle(() => window.shareConfig);
    console.log(await aHandle.jsonValue());

})();
```

## 原生fetch上传
```js
async function doUploadFile(file) {
    var url = window.location.href + file.name
    const formData = new FormData()
    formData.append('file', file)
    return await fetch(url, {
        method: 'PUT',
        body: formData
    })
}
```


## 网页保存文件

```js
savefile = (data, name) => {
    //Blob为js的一个对象，表示一个不可变的, 原始数据的类似文件对象，这是创建文件中不可缺少的！
    var urlObject = window.URL || window.webkitURL || window;
    var export_blob = new Blob([data]);
    var save_link = document.createElementNS("http://www.w3.org/1999/xhtml", "a")
    save_link.href = urlObject.createObjectURL(export_blob);
    save_link.download = name;
    save_link.click();
}

```