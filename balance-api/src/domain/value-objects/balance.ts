/** Value object for Balance */
export class Balance {
    private readonly balance: number;

    constructor(balance: number) {
        this.balance = balance;
    }

    public getValue(): number {
        return this.balance;
    }
}