import {ContactService} from "./contact";
import {Config} from "./config/config";

export class Sdk {
    public contact: ContactService;
    constructor(baseUrl: string) {
        const config = new Config(baseUrl);
        this.contact = new ContactService(config);
    }
}
