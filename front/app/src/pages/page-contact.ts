import {css, html, LitElement} from "lit";
import {customElement, state} from "lit/decorators.js";
import {Sdk} from "@flo/sdk/index.ts";
import "../components/fg-alert/fg-alert.ts";
import  "wired-elements"
import {FormNewContact} from "../components/fg-form-contact/form-new-contact.ts";

@customElement('page-contact')
export class PageContact extends LitElement {
    @state()
    _form = new FormNewContact()

    sdk: Sdk

    constructor() {
        super();
        this.sdk = new Sdk('http://localhost:8080');
    }

    renderForm(): unknown
    {
        return html`
            <form @submit="${this._submitForm}">
                ${this.form.hasError() ? html`<fg-alert @close-alert="${this._handleCloseAlert}">Erreur pendant l'envoie du formulaire</fg-alert>`: html``}
                <div>
                    <label for="name">Nom</label>
                    ${this.form.errors?.name ? html`<div class="error">${this.form.errors.name}</div>`: html`` }
                    <wired-input placeholder="Enter name" name="name" .value="${this.form.message.name}" @input="${this._handleInput}"></wired-input>    
                </div>
                <div>
                    <label for="email">Email</label>
                    ${this.form.errors?.email ? html`<div class="error">${this.form.errors.email}</div>`: html`` }
                    <wired-input placeholder="Enter email" type="text" name="email" .value="${this.form.message.email}" @input="${this._handleInput}"></wired-input>    
                </div>
                <div>
                    <label for="message">Message</label>
                    ${this.form.errors?.message ? html`<div class="error">${this.form.errors.message}</div>`: html`` }
                    <wired-textarea placeholder="Enter message" name="message" rows="6" .value="${this.form.message.message}" @input="${this._handleMessage}"></wired-textarea>    
                </div>
                <div class="actions">
                    <wired-button elevation="2" ?disabled=${!this.form.isValid() || this.form.isLoading()} @click="${this._submitForm}">
                    ${this.form.isLoading() ? html`Ca envoie ;)` : html `Envoyer`}
                    </wired-button>
                </div>
            </form>
        `
    }

    render(): unknown {
        return html `
            <wired-card elevation="5">
                <h4>Contact</h4>
                ${this.form.isSent() ? html`
                    <div>
                        <h3>Message envoyé.</h3>
                        <wired-link @click="${this._handleRenewForm}">Envoyé à nouveau</wired-link>
                    </div>
                ` : this.renderForm()}
                
            </wired-card>
        `
    }

    set form(value) {
        const oldValue = this._form;
        this._form = value;
        this.requestUpdate('form', oldValue); // Force la mise à jour
    }

    get form() {
        return this._form;
    }

    async _submitForm(e: Event): Promise<void> {
        e.preventDefault();
        if (!this.form.isValid() || this.form.isLoading()) {
            return Promise.resolve();
        }

        this.form.resetErrors();
        this.form.loading = true;
        this.requestUpdate();

        try {
            const r = await this.sdk.contact.postANewMessage(this.form.message);
            if (r.status >= 400 && r.status < 500) {
                const data = await r.json();
                this.form.errors = {
                    name: data.errors?.name,
                    email: data.errors?.email,
                    message: data.errors?.message
                };
                this.form.loading = false;
                this.requestUpdate();
                return;
            }
        } catch (err) {
            // @ts-ignore
            if (err?.message === 'Failed to fetch') {
                this.form.errorNetworkOnSubmit = true;
            }
            this.form.loading = false;
            this.requestUpdate();
            return;
        }
        this.form.loading = false;
        this.form.sent = true;
        this.requestUpdate();
    }

    _handleInput(event: Event) {
        // @ts-ignore
        const { name, value } = event.target;
        if(!this.form.message.hasOwnProperty(name)) {
            return;
        }
        // @ts-ignore
        this.form.message[name] = value;
        this.requestUpdate();
    }

    _handleMessage(event: Event) {
        // @ts-ignore
        const { value } = event.target;

        // @ts-ignore
        this.form.message.message = value;
        this.requestUpdate();
    }

    _handleCloseAlert() {
        this.form.resetErrors();
        this.requestUpdate();
    }

    _handleRenewForm(e: Event)
    {
        e.preventDefault();
        this.form.reset();
        this.requestUpdate();
        this.render();
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
        'form-contact': PageContact
    }
}