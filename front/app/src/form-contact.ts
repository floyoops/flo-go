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
    sent: boolean

    @state()
    showErrorSubmit: boolean;

    @state()
    errorsForm: {name?: string, email?: string, message?: string} | null

    sdk: Sdk

    constructor() {
        super();
        this.sdk = new Sdk('http://localhost:8080');
        this.dto = NewMessageDto.fromEmpty();
        this.sent = false;
        this.isLoading = false;
        this.isDisabled = true;
        this.showErrorSubmit = false;
        this.errorsForm = null;
    }

    renderForm(): unknown
    {
        return html`
            <form @submit="${this._submitForm}">
                ${this.showErrorSubmit ? html`<fg-alert @close-alert="${this._handleCloseAlert}">Erreur pendant l'envoie du formulaire</fg-alert>`: html``}
                <div>
                    <label for="name">Nom</label>
                    ${this.errorsForm?.name ? html`<div class="error">${this.errorsForm.name}</div>`: html`` }
                    <wired-input placeholder="Enter name" name="name" .value="${this.dto.name}" @input="${this._handleInput}"></wired-input>    
                </div>
                <div>
                    <label for="email">Email</label>
                    ${this.errorsForm?.email ? html`<div class="error">${this.errorsForm.email}</div>`: html`` }
                    <wired-input placeholder="Enter email" type="text" name="email" .value="${this.dto.email}" @input="${this._handleInput}"></wired-input>    
                </div>
                <div>
                    <label for="message">Message</label>
                    ${this.errorsForm?.message ? html`<div class="error">${this.errorsForm.message}</div>`: html`` }
                    <wired-textarea placeholder="Enter message" name="message" rows="6" .value="${this.dto.message}" @input="${this._handleMessage}"></wired-textarea>    
                </div>
                <div class="actions">
                    <wired-button elevation="2" ?disabled=${this.isDisabled} @click="${this._submitForm}">
                    ${this.isLoading ? html`Ca envoie ;)` : html `Envoyer`}
                    </wired-button>
                </div>
            </form>
        `
    }

    render(): unknown {
        return html `
            <wired-card elevation="5">
                <h4>Contact</h4>
                ${this.sent ? html`<div>Message envoyé. <wired-link @click="${this._handleRenewForm()}">Envoyé à nouveau</wired-link></div>` : this.renderForm()}
                
            </wired-card>
        `
    }

    _initForm() {
        this.isLoading = false;
        this.isDisabled = false;
        this.showErrorSubmit = false;
        this.errorsForm = null;
        this.dto = new NewMessageDto('', '', '');
    }

    async _submitForm(e: Event): Promise<void> {
        e.preventDefault();
        this.errorsForm = null;
        if (this.isDisabled) {
            return Promise.resolve();
        }
        if (!this.dto.isValid()) {
            return;
        }
        this.isLoading = true;
        this.isDisabled = true;
        try {
            const r = await this.sdk.contact.postANewMessage(this.dto);
            if (r.status >= 400 && r.status < 500) {
                const data = await r.json();
                this.errorsForm = {
                    name: data.errors?.name,
                    email: data.errors?.email,
                    message: data.errors?.message
                };
                this.showErrorSubmit = true;
            }
            if (r.status === 201) {
                this._initForm();
            }
        } catch (e) {
            this.showErrorSubmit = true;
        }

        this.isLoading = false;
        this.isDisabled = false;
        this.sent = true;
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
        this.errorsForm = null;
        this.isDisabled = !this.dto.isValid();
    }

    _handleMessage(event: Event) {
        // @ts-ignore
        const { value } = event.target;

        // @ts-ignore
        this.dto['message'] = value;
        this.showErrorSubmit = false;
        this.isDisabled = !this.dto.isValid();
    }

    _handleCloseAlert() {
        this.showErrorSubmit = false;
    }

    _handleRenewForm()
    {
        this.sent = false;
    }

    static styles = [css`
        form > div {
            margin-bottom: 20px;
        }
        .error {
            color: red;
            text-decoration: underline
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