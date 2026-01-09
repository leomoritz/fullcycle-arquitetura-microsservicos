import { EventMessage } from "../messages/message";

/**
 * Interface for producing events to a messaging system.
 */
export interface IEventProducer {
  /**
   * Produces an event to the messaging system.
   * @param input - The event data to be sent.
   * @returns A promise that resolves when the event has been successfully produced.
   */
  publish(input: EventMessage): Promise<void>;
}