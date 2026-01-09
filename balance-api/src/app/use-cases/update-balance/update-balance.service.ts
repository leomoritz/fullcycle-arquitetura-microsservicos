/** Service for updating account balance (implements UpdateBalanceUseCase) */

import { Injectable, Inject } from '@nestjs/common';
import { IUpdateBalanceUseCase } from '../../ports/in/update-balance.usecase';
import { UpdateBalanceInputDto } from 'src/app/dto/input/update-balance-input.dto';
import * as repository from 'src/app/ports/out/accounts-balance.gateway';
import { AccountMapper } from 'src/app/mappers/account.mapper';
import { Account } from 'src/domain/entities/account';
import * as publisher from 'src/app/ports/out/event.publisher.gateway';

@Injectable()
export class UpdateBalanceService implements IUpdateBalanceUseCase {
  constructor(
    @Inject('IAccountsBalanceGateway')
    private readonly accountsBalanceGateway: repository.IAccountsBalanceGateway,
    @Inject('IEventPublisherGateway')
    private readonly eventPublisherGateway: publisher.IEventPublisherGateway,
  ) { }

  async execute(input: UpdateBalanceInputDto): Promise<void> {
    try {
      const account: Account = AccountMapper.toAccountEntity(input);
      account.validate();

      const accountId = account.getAccountId().getValue();
      const newBalance = account.getBalance().getValue();
      const previousBalance = await this.accountsBalanceGateway.getBalanceByAccountId(accountId);
      const balanceUpdatedEvent = AccountMapper.toUpdateBalanceEvent(accountId, previousBalance, newBalance);
      
      await this.accountsBalanceGateway.saveBalance(account);
      this.eventPublisherGateway.publish(balanceUpdatedEvent);
    } catch (error) {
      const failMessage = error instanceof Error ? error.message : 'Unknown error';
      throw new Error(`Failed to update balance: ${failMessage}`);
    }
  }
}