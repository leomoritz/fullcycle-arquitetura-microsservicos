/**
 * Wallet Event Consumer
 * Handles balance updated events from wallet-core via Kafka
 */
import { Logger, Inject, Controller } from '@nestjs/common';
import { Ctx, EventPattern, KafkaContext, Payload } from '@nestjs/microservices';
import * as usecase from 'src/app/ports/in/update-balance.usecase';
import { WalletUpdateBalanceEventData } from './wallet.payload.data';
import { AccountMapper } from 'src/app/mappers/account.mapper';

interface KafkaMessageEnvelope {
  name: string;
  payload: WalletUpdateBalanceEventData;
}

@Controller()
export class WalletEventConsumer {
  private readonly logger = new Logger(WalletEventConsumer.name);

  constructor(
    @Inject('IUpdateBalanceUseCase')
    private readonly updateBalanceUseCase: usecase.IUpdateBalanceUseCase,
  ) { }

  /**
   * Handles wallet balance updated events
   * Topic: balances
   * @param message Event message containing payload with balance updates
   * @param context Kafka context
   */
  @EventPattern('balances')
  async handleBalanceUpdated(@Payload() message: any, @Ctx() context: KafkaContext): Promise<void> {
    this.logEventReceived(context);

    if (!this.validateMessageStructure(message)) {
      this.logger.warn('Invalid message structure received in WalletEventConsumer');
      return;
    }

    try {
      const eventData = message as KafkaMessageEnvelope;

      this.logger.log(`Event name: ${eventData.name}`);
      this.logger.log(`Payload: ${JSON.stringify(eventData.payload)}`);

      const payload = eventData.payload;

      const accountFromInputDto = AccountMapper.toUpdateBalanceInputDto(
        payload.account_id_from,
        payload.balance_account_id_from
      );

      const accountToInputDto = AccountMapper.toUpdateBalanceInputDto(
        payload.account_id_to,
        payload.balance_account_id_to
      );

      await this.updateBalanceUseCase.execute(accountFromInputDto);
      await this.updateBalanceUseCase.execute(accountToInputDto);

      this.logger.log(`Successfully updated balances for accounts:
        ${payload.account_id_from} and ${payload.account_id_to}`);
    } catch (error) {
      this.logger.error( //
        `Error processing balance update event`, //
        error.stack, //
        JSON.stringify(message)
      );
      throw error;
    }
  }

  private logEventReceived(context: KafkaContext) {
    const rawMessage = context.getMessage();
    const topicName = context.getTopic();
    this.logger.log(`Received message from topic ${topicName}, partition ${context.getPartition()}, offset ${rawMessage.offset}`);
  }

  private validateMessageStructure(message: any): boolean {
    if (message && typeof message === 'object') {
      return 'name' in message && 'payload' in message;
    }
    return false;
  }
}
