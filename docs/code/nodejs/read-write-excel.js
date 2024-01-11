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