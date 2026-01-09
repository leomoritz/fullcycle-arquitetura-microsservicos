/** Event publisher interface */

import { BalanceUpdatedEvent } from "src/domain/events/balance-updated.event";

export interface IEventPublisherGateway {

  /**
   * Method to publish an event
   * @param event Event to be published
   */
  publish(event: BalanceUpdatedEvent): void;
}