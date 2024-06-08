import {Config} from "../config/config";
import {NewMessageDto} from "./new-message.dto";

export class ContactService {

    private headers = {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
    };

    constructor(private readonly config: Config) {
    }

    public async postANewMessage(data: NewMessageDto): Promise<Response> {
        const url = `${this.config.baseUrl}/post-a-new-message-contact`;
        return fetch(url, {
            method: 'POST',
            headers: this.headers,
            body: data.toJsonStringify(),
        });
    }
}
