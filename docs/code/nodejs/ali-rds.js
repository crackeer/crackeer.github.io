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