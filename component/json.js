const myTemplate = document.createElement('template')
myTemplate.innerHTML = `
<div id="jsoneditor" style="width: 100%; height: calc(100vh - 100px);"></div>
`

const myTemplate = document.createElement('script')
myTemplate.innerHTML = `
<div id="jsoneditor" style="width: 100%; height: calc(100vh - 100px);"></div>
`
import 
class JSON extends HTMLElement {
    static get observedAttributes() {
        return ['height', 'icon'];
    }

    constructor() {
        super();
        console.log(this.getAttribute('height'));
        this._shadowRoot = this.attachShadow({ mode: 'open' })
        let node = myTemplate.content.cloneNode(true)
        node.style ={height:this.getAttribute('height')}
        const myTemplate = document.createElement('script')

        this._shadowRoot.appendChild(node)
    }
    render() {
        console.log("render")
        this._shadowRoot = this.attachShadow({ mode: 'open' })
        
        this._shadowRoot.appendChild(myTemplate.content.cloneNode(true))
      }
    // More JS goes here
}

window.customElements.define('custom-json', JSON)

