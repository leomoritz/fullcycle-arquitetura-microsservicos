/** Account Repository (implements AccountRepository) */
import { Injectable } from '@nestjs/common';
import { Repository } from 'typeorm';
import { InjectRepository } from '@nestjs/typeorm';
import { AccountEntity } from '../entities/account.entity';
import { Account } from 'src/domain/entities/account';
import { AccountMapper } from 'src/app/mappers/account.mapper';
import { IAccountsBalanceGateway } from 'src/app/ports/out/accounts-balance.gateway';

@Injectable()
export class AccountRepository implements IAccountsBalanceGateway {
    constructor(
        @InjectRepository(AccountEntity)
        private readonly accountRepo: Repository<AccountEntity>,
    ) { }

    /**
     * Save account's balance.
     * If account exists, update balance; otherwise, create new account.
     * @param account data to be saved
     */
    async saveBalance(account: Account): Promise<void> {
        const accountEntity = AccountMapper.toDBEntity(account);
        await this.accountRepo.upsert(accountEntity, ['accountId']);
    }

    /**
   * Get Account's balance by account id.
   * @param accountId account's identifier
   * @returns current balance from account.
   */
    async getBalanceByAccountId(accountId: string): Promise<number> {
        const account = await this.accountRepo.findOne({ where: { accountId } });
        return account ? account.balance : 0;
    }
}