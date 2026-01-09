/**
 * Kafka Event Producer
 */
import { Injectable, Logger } from '@nestjs/common';
import { Kafka, Producer, ProducerRecord } from 'kafkajs';
import { IEventProducer } from './event.producer';
import { EventMessage } from '../messages/message';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class KafkaProducer implements IEventProducer {
  private readonly logger = new Logger(KafkaProducer.name);
  private readonly configService: ConfigService;
  private kafka: Kafka;
  private producer: Producer;

  constructor(configService: ConfigService) {
    this.configService = configService;
    
    this.kafka = new Kafka({
      clientId: this.configService.get<string>('KAFKA_CLIENT_ID') || 'balance-api',
      brokers: [this.configService.get<string>('KAFKA_BROKER') || 'kafka:29092'],
    });
    this.producer = this.kafka.producer();
    this.connect();
  }

  private async connect() {
    await this.producer.connect();
    this.logger.log('Kafka producer connected');
  }

  public async publish(input: EventMessage) {
    const record: ProducerRecord = {
      topic: input.topic,
      messages: input.message,
    };

    try {
      await this.producer.send(record);
      this.logger.log(`Event sent to topic ${input.topic}`);
    } catch (error) {
      this.logger.error(`Failed to send event to topic ${input.topic}`, error.stack);
    }
  }
}