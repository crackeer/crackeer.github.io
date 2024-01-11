/// [ali-rds]
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
db.delete('table_name', {
    id : 1022
})
/// [ali-rds]


/// [read-file]
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
/// [read-file]


/// [excel-lib]
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

// 写Excel
var writeExcel = async () => {
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
/// [excel-lib]


/// [puppeteer-crawl]
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
/// [puppeteer-crawl]

/// [js-save-file]
var savefile = (data, name) => {
    //Blob为js的一个对象，表示一个不可变的, 原始数据的类似文件对象，这是创建文件中不可缺少的！
    var urlObject = window.URL || window.webkitURL || window;
    var export_blob = new Blob([data]);
    var save_link = document.createElementNS("http://www.w3.org/1999/xhtml", "a")
    save_link.href = urlObject.createObjectURL(export_blob);
    save_link.download = name;
    save_link.click();
}
/// [js-save-file]


/// [fetch]
async function doUploadFile(file) {
    var url = window.location.href + file.name
    const formData = new FormData()
    formData.append('file', file)
    return await fetch(url, {
        method: 'PUT',
        body: formData
    })
}
/// [fetch]