/** Event triggered when an account balance is updated */

export class BalanceUpdatedEvent {
  readonly accountId: string;
  readonly previousBalance: number;
  readonly newBalance: number;
  readonly timestamp: Date;

    constructor(props: {
    accountId: string;
    previousBalance: number;
    newBalance: number;
    timestamp?: Date;
  }) {
    this.accountId = props.accountId;
    this.previousBalance = props.previousBalance;
    this.newBalance = props.newBalance;
    this.timestamp = props.timestamp || new Date();
  }
}