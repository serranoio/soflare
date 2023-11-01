import { CSSResultGroup, LitElement, css, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'
import { TABS } from '../../model/vars';

@customElement('nav-component')
export class Main extends LitElement {

    @property()
    direction = "";

    static styles?: CSSResultGroup | undefined = css`
        nav {
            height: 80px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-size: 2rem;
        }

        * {
        padding: 0;
        margin: 0;
        box-sizing: border-box;
      }
      
        ul {
            display: flex;
            justify-content: space-between;
            align-items: center;
            list-style: none;
            gap: 2rem;
        }
        a {
            
            text-decoration: none;
        }

    `

  @property()
  title: string = "";

  render() {
    return html`
    <nav>
    <h3>${this.title}</h3>
    <ul style="flex-direction: ${this.direction};">
        ${TABS.map((tab: string) => {
            return html`<li><a href=${`#${tab.toLowerCase()}`}>${tab}</a></li>`
        })}
    </ul>
    </nav>
    `
  }
}
