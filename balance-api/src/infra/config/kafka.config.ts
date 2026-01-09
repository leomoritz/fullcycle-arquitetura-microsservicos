/**
 * Kafka configuration - Centralized Kafka instance
 */
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class KafkaConfig {
  constructor(private configService: ConfigService) {}

  getKafkaConfig() {
    return {
      brokers: this.configService.get('KAFKA_BROKER').split(','),
    };
  }
}