function getQuery(key, value) {
    let url = new URLSearchParams(window.location.search)
    return url.get(key) || value
}

function pushStateWith(query) {
    let newURL = window.location.pathname + "?" + httpBuildQuery(query)
    window.history.pushState(null, "", newURL)
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

function httpBuildQuery(query){
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

function parseBookmark(content) {
    let parts = content.split("\n")
    let list = []
    console.log(parts)
    for(var i in parts) {
        if(parts[i].length > 0 ) {
            let temp = parts[i].split(">")
            if(temp.length == 1) {
                list.push({
                    title : temp[0],
                    href : temp[0]
                })
            } else if(temp.length > 1) {
                list.push({
                    title : temp[1],
                    href : temp[0]
                })
            }
        }
    }
    return list
 }