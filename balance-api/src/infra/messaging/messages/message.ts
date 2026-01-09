/**
 * Class to define event publisher behavior
 */

class Message {
    key: string | undefined;
    value: any;
    headers?: Record<string, any>;
}

export class EventMessage {
    topic: string;
    message: Message[];
}