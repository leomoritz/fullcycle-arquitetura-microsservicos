import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { INestApplication } from '@nestjs/common';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService);
  const { kafkaBroker, kafkaGroupId } = await configureKafkaMicroserviceAndConnect(configService, app);

  const port = configService.get('APP_PORT') ?? 3003;
  await app.listen(port);

  console.log(`Balance API running on port ${port}`);
  console.log(`Kafka consumer active - listening to configured topics`);
  console.log(`Kafka Config: Broker=${kafkaBroker}, GroupId=${kafkaGroupId}`);
}

bootstrap().catch((error) => {
  console.error('Failed to start application:', error);
  process.exit(1);
});


async function configureKafkaMicroserviceAndConnect(configService: ConfigService<unknown, boolean>, app: INestApplication<any>) {
  const kafkaBroker = configService.get<string>('KAFKA_BROKER');
  const kafkaGroupId = configService.get<string>('KAFKA_GROUP_ID');

  if (!kafkaBroker || !kafkaGroupId) {
    throw new Error(
      'Missing required environment variables: KAFKA_BROKER and/or KAFKA_GROUP_ID'
    );
  }

  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.KAFKA,
    options: {
      client: {
        clientId: 'balance-api',
        brokers: [kafkaBroker],
        connectionTimeout: 10000,
        requestTimeout: 30000,
      },
      consumer: {
        groupId: kafkaGroupId,
        allowAutoTopicCreation: true,
        sessionTimeout: 30000,
        heartbeatInterval: 3000,
      },
      subscribe: {
        fromBeginning: true,
      },
      run: {
        autoCommit: true,
      },
    }
  });

  await app.startAllMicroservices();
  return { kafkaBroker, kafkaGroupId };
}

