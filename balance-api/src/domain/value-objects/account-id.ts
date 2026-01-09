/** Value object for Account ID */
export class AccountId {
  private readonly id: string;

  constructor(id: string) {
    this.id = id;
  }

  public getValue(): string {
    return this.id;
  }
}