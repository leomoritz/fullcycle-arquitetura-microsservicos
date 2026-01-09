/** DTO for getting account balance */
export class GetBalanceOutputDto {
  accountId: string;
  balance: number;

    constructor(accountId: string, balance: number) {
        this.accountId = accountId;
        this.balance = balance;
    }
}