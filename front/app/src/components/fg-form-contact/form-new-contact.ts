import {NewMessageDto} from "@flo/sdk/index.ts";

export class FormNewContact
{
    public sent: boolean;
    public errors: {name?: string, email?: string, message?: string} | null
    public loading: boolean;
    public errorNetworkOnSubmit: boolean;
    public message: NewMessageDto;

    constructor() {
        this.sent = false;
        this.errors = null;
        this.loading = false;
        this.errorNetworkOnSubmit = false;
        this.message = NewMessageDto.fromEmpty();
    }

    public hasError(): boolean {
        return this.errors !== null || this.errorNetworkOnSubmit;
    }

    public isLoading(): boolean {
        return this.loading;
    }

    public isSent(): boolean {
        return this.sent;
    }

    public resetErrors(): void {
        this.errors = null;
        this.errorNetworkOnSubmit = false;
    }

    public reset(): void {
        this.resetErrors();
        this.sent = false;
        this.message = NewMessageDto.fromEmpty();
    }

    public isValid(): boolean {
        return this.message.isValid();
    }
}
