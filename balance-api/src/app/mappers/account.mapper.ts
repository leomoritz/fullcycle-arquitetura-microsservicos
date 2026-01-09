import { Account } from "src/domain/entities/account";
import { GetBalanceOutputDto } from "../dto/output/get-balance-output.dto";
import { UpdateBalanceInputDto } from "../dto/input/update-balance-input.dto";
import { AccountId } from "src/domain/value-objects/account-id";
import { BalanceUpdatedEvent } from "src/domain/events/balance-updated.event";
import { AccountEntity } from "src/infra/database/entities/account.entity";

/** Account mapper */
export class AccountMapper {

  public static toGetBalanceOutputDTO(accountId: string, balance: number): GetBalanceOutputDto {
    return new GetBalanceOutputDto(accountId, balance);
  }

  public static toAccountEntity(input: UpdateBalanceInputDto): Account {
    return new Account(new AccountId(input.accountId), input.balance);
  }

  public static toUpdateBalanceEvent(accountId: string, previousBalance: number, newBalance: number): BalanceUpdatedEvent {
    return new BalanceUpdatedEvent({accountId, previousBalance, newBalance});
  }

  public static toDBEntity(account: Account): AccountEntity {
    const accountDB = new AccountEntity();
    accountDB.accountId = account.getAccountId().getValue();
    accountDB.balance = account.getBalance().getValue();
    return accountDB;
  }

  public static toUpdateBalanceInputDto(accountId: string, balance: number): UpdateBalanceInputDto {
    return new UpdateBalanceInputDto(accountId, balance);
  }
}