export class EventListener {
    handler: any
    once: boolean

    constructor(handler: any, once: boolean) {
        this.handler = handler
        this.once = once
    }
}
