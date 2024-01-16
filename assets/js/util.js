function getQuery(key, value) {
    let url = new URLSearchParams(window.location.search)
    return url.get(key) || value
}

function pushStateWith(query) {
    let newURL = window.location.pathname + "?" + httpBuildQuery(query)
    window.history.pushState(null, "", newURL)
}

function nanoid(t) {
    return crypto.getRandomValues(new Uint8Array(t)).reduce(((t, e) => t += (e &= 63) < 36 ? e.toString(36) : e < 62 ? (e - 26).toString(36).toUpperCase() : e > 62 ? "-" : "_"), "")
}

function simpleReload() {
    window.location.reload()
}

function reloadWith(query) {
    window.location.href = window.location.pathname + '?' + httpBuildQuery(query)
}

function jump(path, query) {
    window.location.href = path + '?' + httpBuildQuery(query)
}

function buildQuery(query) {
    let params = new URLSearchParams("")
    Object.keys(query).forEach(k => {
        params.append(k, query[k])
    })
    return params.toString()
}

function cloneObject(data) {
    let raws = JSON.stringify(data)
    return JSON.parse(raws)
}


function saveFile(data, name) {
    //Blob为js的一个对象，表示一个不可变的, 原始数据的类似文件对象，这是创建文件中不可缺少的！
    var urlObject = window.URL || window.webkitURL || window;
    var export_blob = new Blob([data]);
    var save_link = document.createElementNS("http://www.w3.org/1999/xhtml", "a")
    save_link.href = urlObject.createObjectURL(export_blob);
    save_link.download = name;
    save_link.click();
}

//js 读取文件
function readFiles(file) {
    var reader = new FileReader();//new一个FileReader实例
    if (/text+/.test(file.type)) {//判断文件类型，是不是text类型
        reader.onload = function (result) {
            console.log(result)
        }
        reader.readAsText(file);
    } else if (/image+/.test(file.type)) {//判断文件是不是imgage类型
        reader.onload = function (result) {
            console.log(result)
        }
        reader.readAsDataURL(file);
    }
    
}
