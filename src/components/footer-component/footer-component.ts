import { CSSResultGroup, LitElement, css, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'

@customElement('footer-component')
export class FooterComponent extends LitElement {

  static styles?: CSSResultGroup | undefined = css`
    * {
    padding: 0;
    margin: 0;
    box-sizing: border-box;
  }

  :host {
      max-width: 140rem;
      display: block;
      margin: 0 auto;

  }
  
    footer {
        border-top: 1px solid black;
        padding: 9.8rem 0;
        display: grid;
        grid-template-columns: 1fr 1fr;
    }
    ul {
        list-style: none;
        display: flex;
        flex-direction: column;
    }

    `

  @property()
  title: string = "";

  render() {
    return html`
    <footer>
        <div class="contact">
            <ul>
            <li>Cool Person</li>
            </ul>
        </div>
        <slot name="nav"></slot>
    </footer>
    `
  }
}
