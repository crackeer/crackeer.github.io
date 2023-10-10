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
