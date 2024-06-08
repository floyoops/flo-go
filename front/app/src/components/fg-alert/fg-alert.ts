import {css, html, LitElement} from "lit";
import {customElement} from "lit/decorators.js";

@customElement('fg-alert')
export class FgAlert extends LitElement {
    render() {
        return html`
            <div class="alert">
                <span class="closebtn" @click=${this._handleClose}>&times;</span>
                <slot></slot>
            </div>
        `;
    }

    _handleClose() {
        const event = new CustomEvent('close-alert', {
            bubbles: true,
            composed: true
        });
        this.dispatchEvent(event);
    }

    static styles = css`
        .alert {
            padding: 5px;
            background-color: firebrick; /* Red */
            color: white;
            margin-bottom: 15px;
            position: relative;
        }
        .closebtn {
            margin-left: 15px;
            color: white;
            font-weight: bold;
            float: right;
            font-size: 22px;
            line-height: 20px;
            cursor: pointer;
            transition: 0.3s;
        }
        .closebtn:hover {
            color: black;
        }
  `;
}
