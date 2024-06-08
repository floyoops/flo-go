export class NewMessageDto  {
    constructor(
        public name: string,
        public email: string,
        public message: string
    ) {
    }

    toJsonStringify(): string
    {
        return JSON.stringify(this);
    }

    static fromEmpty() : NewMessageDto {
        return new NewMessageDto('', '', '');
    }

    isValid(): boolean
    {
        if (this.name.length < 2) {
            return false;
        }
        if (this.email.length < 4) {
            return false;
        }
        if (this.message.length < 2) {
            return false;
        }
        return true;
    }
}
