import {css, html, LitElement} from "lit";
import {customElement, state} from "lit/decorators.js";
import {NewMessageDto, Sdk} from "@flo/sdk";
import "./components/fg-alert/fg-alert.ts";
import  "wired-elements"

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
        return html `
            <wired-card elevation="5">
                <h4>Contact</h4>
                <form @submit="${this._submitForm}">
                    ${this.showErrorSubmit ? html`<fg-alert @close-alert="${this._handleCloseAlert}">Erreur pendant l'envoie du formulaire</fg-alert>`: html``}
                    <div>
                        <label for="name">Nom</label>
                        <wired-input placeholder="Enter name" name="name" .value="${this.dto.name}" @input="${this._handleInput}"></wired-input>    
                    </div>
                    <div>
                        <label for="email">Email</label>
                        <wired-input placeholder="Enter email" type="text" name="email" .value="${this.dto.email}" @input="${this._handleInput}"></wired-input>    
                    </div>
                    <div>
                        <label for="message">Message</label>
                        <wired-textarea placeholder="Enter message" name="message" rows="6" .value="${this.dto.message}" @input="${this._handleMessage}"></wired-textarea>    
                    </div>
                    <div class="actions">
                        <wired-button elevation="2" ?disabled=${this.isDisabled} @click="${this._submitForm}">
                        ${this.isLoading ? html`Ca envoie ;)` : html `Envoyer`}
                        </wired-button>
                    </div>
                </form>
            </wired-card>
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
        if (this.isDisabled) {
            return Promise.resolve();
        }
        console.log('_submitForm', this.dto);
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

    _handleMessage(event: Event) {
        const { value } = event.target;
        console.log('_handleMessage', value);

        // @ts-ignore
        this.dto['message'] = value;
        this.showErrorSubmit = false;
        this.isDisabled = !this.dto.isValid();
    }

    _handleCloseAlert() {
        this.showErrorSubmit = false;
    }

    static styles = [css`
        form > div {
            margin-bottom: 20px;
        }
        .actions {
            text-align: "center";
        }
        label {
            display: block;
        }
        wired-card {
            padding: 20px;
        }
        wired-input {
            width: 200px;
            font-family: inherit;
        }
        wired-textarea {
            font-family: inherit;
        }
    `];
}

declare global {
    interface HTMLElementTagNameMap {
        'form-contact': FormContact
    }
}