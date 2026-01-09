import { DomainException } from "src/shared/exceptions/domain.exception";
import { AccountId } from "../value-objects/account-id";
import { Balance } from "../value-objects/balance";

/** Account domain entity */
export class Account {
    private _accountId: AccountId;
    private _balance: Balance;

    constructor(accountId: AccountId, balance: number) {
        this._accountId = accountId;
        this.updateBalance(balance);
    }

    public getAccountId(): AccountId {
        return this._accountId;
    }

    public getBalance(): Balance {
        return this._balance;
    }

    private updateBalance(balance: number): void {
        this._balance = new Balance(balance);
    }

    public validate(): void {
        if (!this._accountId || !this._accountId.getValue()) {
            throw new DomainException("Invalid account ID");
        }
        if (this._balance.getValue() < 0) {
            throw new DomainException("Balance cannot be negative");
        }
    }
}