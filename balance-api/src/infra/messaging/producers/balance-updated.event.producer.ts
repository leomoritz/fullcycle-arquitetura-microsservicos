/** 
 * Producer to publish BalanceUpdatedEvent to Kafka
 */

import { Injectable, Logger, Inject } from '@nestjs/common';
import { BalanceUpdatedEvent } from 'src/domain/events/balance-updated.event';
import { IEventPublisherGateway } from 'src/app/ports/out/event.publisher.gateway';
import { KafkaProducer } from './kafka.producer';
import { EventMessage } from '../messages/message';

@Injectable()
export class BalanceUpdatedEventProducer implements IEventPublisherGateway {
    private readonly logger = new Logger(BalanceUpdatedEventProducer.name);
    private readonly kafkaProducer: KafkaProducer;

    constructor(kafkaProducer: KafkaProducer) {
        this.kafkaProducer = kafkaProducer;
    }

    publish(event: BalanceUpdatedEvent): void {
        const envelope = {
            name: 'BalanceUpdated',
            payload: {
                account_id: event.accountId,
                previous_balance: event.previousBalance,
                new_balance: event.newBalance,
                timestamp: event.timestamp.toISOString(),
            }
        };

        const eventMessage = new EventMessage();
        eventMessage.topic = 'currentbalances';
        eventMessage.message = [
            {
                key: event.accountId,
                value: JSON.stringify(envelope),
            }
        ];
        
        this.kafkaProducer.publish(eventMessage);
        this.logger.log(`Published BalanceUpdatedEvent for accountId: ${event.accountId} (${event.previousBalance} → ${event.newBalance})`);
    }
}