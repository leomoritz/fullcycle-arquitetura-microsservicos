/** DTO for updating account balance */
export class UpdateBalanceInputDto {
  accountId: string;
  balance: number;

  constructor(accountId: string, balance: number) {
    this.accountId = accountId;
    this.balance = balance;
  }
}