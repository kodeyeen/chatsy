import { EventListener } from "./listener"

export class EventDispatcher {
    listeners: Record<string, EventListener[]> = {}

    constructor() {
        this.listeners = {}
    }

    dispatch<T>(type: string, event: T) {
        const listeners = this.listeners[type]

        for (const listener of listeners) {
            listener.handler(event)

            if (listener.once) {
                this.off(type, listener.handler)
            }
        }
    }

    on(type: string, handler: any) {
        if (!this.has(type)) {
            this.listeners[type] = []
        }

        const listeners: EventListener[] = this.listeners[type]

        listeners.push(new EventListener(handler, false))

        this.listeners[type] = listeners
    }

    once(type: string, handler: any) {
        const listeners = this.listeners[type]

        listeners.push(new EventListener(handler, true))

        this.listeners[type] = listeners
    }

    off(type: string, handler: any) {
        const listeners: EventListener[] = this.listeners[type]

        const index: number = listeners.indexOf(handler)

        if (index != -1) {
            listeners.splice(index, 1);
        }

        if (listeners.length == 0) {
            delete this.listeners[type]
        } else {
            this.listeners[type] = listeners
        }
    }

    has(type: string): boolean {
        return this.listeners.hasOwnProperty(type)
    }
}
