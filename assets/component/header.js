const myTemplate = document.createElement('template')
myTemplate.innerHTML = `<nav class="navbar navbar-default">
    <div class="container-fluid">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse"
                data-target="#navbar-collapse-1" aria-expanded="false">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="https://github.com/crackeer" target="_blank">
                <svg height="25" aria-hidden="true" viewBox="0 0 16 16" version="1.1" width="25"
                    data-view-component="true" class="octicon octicon-mark-github v-align-middle">
                    <path
                        d="M8 0c4.42 0 8 3.58 8 8a8.013 8.013 0 0 1-5.45 7.59c-.4.08-.55-.17-.55-.38 0-.27.01-1.13.01-2.2 0-.75-.25-1.23-.54-1.48 1.78-.2 3.65-.88 3.65-3.95 0-.88-.31-1.59-.82-2.15.08-.2.36-1.02-.08-2.12 0 0-.67-.22-2.2.82-.64-.18-1.32-.27-2-.27-.68 0-1.36.09-2 .27-1.53-1.03-2.2-.82-2.2-.82-.44 1.1-.16 1.92-.08 2.12-.51.56-.82 1.28-.82 2.15 0 3.06 1.86 3.75 3.64 3.95-.23.2-.44.55-.51 1.07-.46.21-1.61.55-2.33-.66-.15-.24-.6-.83-1.23-.82-.67.01-.27.38.01.53.34.19.73.9.82 1.13.16.45.68 1.31 2.69.94 0 .67.01 1.3.01 1.49 0 .21-.15.45-.55.38A7.995 7.995 0 0 1 0 8c0-4.42 3.58-8 8-8Z">
                    </path>
                </svg>
            </a>
        </div>

        <div class="collapse navbar-collapse" id="navbar-collapse-1">
            <ul class="nav navbar-nav">
                <li><a href="/">首页</a></li>
                <li><a href="/page/bookmark.html">书签</a></li>
                <li><a href="/page/json.html">JSON编辑</a></li>
                <li><a href="/page/markdown/list.html">Markdown</a></li>
            </ul>
            <ul class="nav navbar-nav navbar-right">
                <li><a href="/page/jstool/index.html" target="_blank">JS工具</a></li>
                <li><a href="/doc.html" target="_blank">文档</a></li>
            </ul>
        </div>
    </div>
</nav>
`
class Header extends HTMLElement {
    static get observedAttributes() {
        return ['theme', 'icon'];
    }

    constructor() {
        super();
                this._shadowRoot = this.attachShadow({ mode: 'open' })
        this._shadowRoot.appendChild(myTemplate.content.cloneNode(true))

    }
    render() {
        console.log("render")
        this._shadowRoot = this.attachShadow({ mode: 'open' })
        this._shadowRoot.appendChild(myTemplate.content.cloneNode(true))
      }
    // More JS goes here
}

window.customElements.define('header-menu', Header)

