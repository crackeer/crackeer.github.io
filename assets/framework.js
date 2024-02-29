
window.VueObject = null
window.VueMount = "#app"

const JS_FILES = [
    '/assets/js/vue.global.js',
    '/assets/js/jquery.js',
    '/assets/bootstrap/bootstrap.bundle.min.js',
    '/assets/bootstrap/bootbox.js',
    '/assets/js/axios.min.js',
    '/assets/js/dayjs.min.js',
    '/assets/js/util.js',
    '/assets/js/filereader.js',
    '/assets/jsoneditor/jsoneditor9.8.js'
]

const DIV_APP_HTML = `<div id="app">
    <div class="d-flex justify-content-center pt-5">
        <div class="spinner-border" role="status">
            <span class="visually-hidden">Loading...</span>
        </div>
    </div>
</div>`

function loadJsUrl(url) {
    return new Promise((resolve) => {
        let domScript = document.createElement("script");
        domScript.src = url;
        domScript.onload = domScript.onreadystatechange = function () {
            if (!this.readyState || 'loaded' === this.readyState || 'complete' === this.readyState) {
                resolve()
            }
        }
        document.getElementsByTagName('head')[0].appendChild(domScript);
    });
}

async function initialize() {
    let result = await fetch("/assets/template/head.html")
    let header = await result.text()
    document.getElementsByTagName('head')[0].insertAdjacentHTML('beforeEnd', header)

    document.getElementsByTagName('body')[0].insertAdjacentHTML('beforeEnd', DIV_APP_HTML)

    // load js
    for (var i = 0; i < JS_FILES.length; i++) {
        await loadJsUrl(JS_FILES[i])
        await sleep(1)
    }

    setTimeout(() => {
        if (window.VueObject != null) {
            console.log(window.VueObject)
            window.Vm = mountVueObject(window.VueObject, window.VueMount)
        }
    }, 200)
}

initialize()

function mountVueObject(object, element) {
    if (Vue === undefined) {
        console.log('mountVueObject, Vue not defined');
        return
    }
    var vm = Vue.createApp(object)
    vm.mount(element)
    return vm
}

function sleep(time) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve();
        }, time);
    });
}
