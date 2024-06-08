import {css, html, LitElement} from "lit";
import {customElement, state} from "lit/decorators.js";
import {NewMessageDto, Sdk} from "@flo/sdk";
import "./components/fg-alert/fg-alert.ts";

@customElement('form-contact')
export class FormContact extends LitElement {
    @state()
    dto: NewMessageDto;

    @state()
    isDisabled: boolean;

    @state()
    isLoading: boolean;

    @state()
    showErrorSubmit: boolean;

    sdk: Sdk

    constructor() {
        super();
        this.sdk = new Sdk('http://localhost:8080');
        this.dto = NewMessageDto.fromEmpty();
        this.isLoading = false;
        this.isDisabled = true;
        this.showErrorSubmit = false;
    }

    render(): unknown {
        return html`
            <div class="form-contact">
                ${this.showErrorSubmit ? html`<fg-alert @close-alert="${this._handleCloseAlert}">Erreur pendant l'envoie du formulaire</fg-alert>`: html``}
                <form name="contact" @submit="${this._submitForm}">
                    <label for="name">Nom</label>
                    <input type="text" name="name" .value="${this.dto.name}" @input="${this._handleInput}">
                    <label for="email">Email</label>
                    <input type="text" name="email" .value="${this.dto.email}" @input="${this._handleInput}">
                    <label for="message">Message</label>
                    <textarea name="message" id="message-0-m" cols="30" rows="10" .value="${this.dto.message}" @input="${this._handleInput}"></textarea>
                    <button ?disabled=${this.isDisabled} @click="${this._submitForm}">
                        ${this.isLoading ? html`Ca envoie ;)` : html `Envoyer`}
                    </button>
                </form>
            </div>
        `
    }

    _initForm() {
        this.isLoading = false;
        this.isDisabled = false;
        this.showErrorSubmit = false;
        this.dto = new NewMessageDto('', '', '');
    }

    async _submitForm(e: Event): Promise<void> {
        e.preventDefault();
        if (!this.dto.isValid()) {
            return;
        }
        this.isLoading = true;
        this.isDisabled = true;
        try {
            const r = await this.sdk.contact.postANewMessage(this.dto);
            if (r.status >= 400 && r.status < 500) {
                this.showErrorSubmit = true;
            }
            if (r.status === 201) {
                this._initForm();
            }
            console.log('r', r);
        } catch (e) {
            this.showErrorSubmit = true;
            console.error('error submitForm', e);
        }

        this.isLoading = false;
        this.isDisabled = false;
    }

    _handleInput(event: Event) {
        // @ts-ignore
        const { name, value } = event.target;
        if(!this.dto.hasOwnProperty(name)) {
            return;
        }
        // @ts-ignore
        this.dto[name] = value;
        this.showErrorSubmit = false;
        this.isDisabled = !this.dto.isValid();
    }

    _handleCloseAlert() {
        this.showErrorSubmit = false;
    }

    static styles = [css`
        .form-contact {
            max-width: 500px;
        }
        .form-contact > form > input {
            width: 98%;
        }
        .form-contact > form > textarea {
            width: 98%;
        }
        .form-contact > form > button {
            margin-top: 10px;
            width: 100%;
        }
    `];
}

declare global {
    interface HTMLElementTagNameMap {
        'form-contact': FormContact
    }
}