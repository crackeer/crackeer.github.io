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