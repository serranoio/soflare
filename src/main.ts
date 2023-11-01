import { CSSResultGroup, LitElement, css, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'
import sup from "./assets/sup.jpg"

@customElement('main-component')
export class Main extends LitElement {

  @property()
  heading: string = `Welcome To Title!`;

  @property()
  bio: string = `this is a bio for poopy buttheads`;

  @property()
  imageUrl: string = `rando-place`;

  static styles?: CSSResultGroup | undefined = css`

  * {
    padding: 0;
    margin: 0;
    box-sizing: border-box;
  }
    main {
      max-width: 130rem;
      margin: 0 auto;
      display: grid;
      grid-template-columns: 1fr 1fr;
      min-height: calc(100vh - 80px);
    }

    .title-box {
      display: flex;
      align-items: center;
      justify-content: center;
      gap: 1rem;
      flex-direction: column;
    }

    h2 {
      font-size: 7.2rem;
    }

    p {
      font-size: 2rem;
    }

    .image-box {
      // max-width: 50vw;
      display: flex;
      align-items: center;
      
    }
    img {
      width: 100%;
    }
  `;

  render() {
    return html`
      <slot name="nav"></slot>
      <main>

        <div class="title-box">
          <h2>${this.heading}</h2>
          <p>${this.bio}</p>
        </div>

        <div class="image-box">
          <img src=${sup}/>
        </div>

      </main>
    `
  }


}
