import { Module } from '@nestjs/common';
import { BalanceController } from './presentation/controllers/balance.controller';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { WalletEventConsumer } from './presentation/consumers/wallet-event.consumer';
import { AccountEntity } from './infra/database/entities/account.entity';
import { AccountRepository } from './infra/database/repositories/account.repository';
import { KafkaConfig } from './infra/config/kafka.config';
import { GetBalanceService } from './app/use-cases/get-balance/get-balance.service';
import { UpdateBalanceService } from './app/use-cases/update-balance/update-balance.service';
import { BalanceUpdatedEventProducer } from './infra/messaging/producers/balance-updated.event.producer';
import { KafkaProducer } from './infra/messaging/producers/kafka.producer';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    TypeOrmModule.forRootAsync({
      useFactory: (config: ConfigService) => ({
        type: 'mysql',
        host: config.get('DB_HOST'),
        port: +config.get('DB_PORT'), // the '+' converts string to number
        username: config.get('DB_USER'),
        password: config.get('DB_PASSWORD'),
        database: config.get('DB_NAME'),
        entities: [AccountEntity],
        synchronize: true, // Note: set to false in production
      }),
      inject: [ConfigService],
    }),
    TypeOrmModule.forFeature([AccountEntity]),
  ],
  controllers: [
    BalanceController, // Controller for balance endpoints
    WalletEventConsumer, // Controller for Kafka wallet events
  ],
  providers: [
    KafkaConfig,
    KafkaProducer,
    // Use Cases
    {
      provide: 'IGetBalanceUseCase',
      useClass: GetBalanceService,
    },
    {
      provide: 'IUpdateBalanceUseCase',
      useClass: UpdateBalanceService,
    },
    // Gateways / Repositories
    {
      provide: 'IAccountsBalanceGateway',
      useClass: AccountRepository,
    },
    {
      provide: 'IEventPublisherGateway',
      useClass: BalanceUpdatedEventProducer,
    },
  ],
})
export class AppModule {}
