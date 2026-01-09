/** Interface for Account's balance Port to persist data */

import { Account } from "src/domain/entities/account";

export interface IAccountsBalanceGateway {
  /**
   * Method to save the balance of an account
   * @param account  Account entity to be saved
   */
  saveBalance(account: Account): Promise<void>;

  /**
   * Method to get the balance of an account by its ID
   * @param accountId ID of the account
   */
  getBalanceByAccountId(accountId: string): Promise<number>;

}
