/** Interface for the GetBalance use case */

import { GetBalanceOutputDto } from "src/app/dto/output/get-balance-output.dto";

export interface IGetBalanceUseCase {

  /**
   * Executes the use case to get the balance of an account
   * @param accountId ID of the account
   * @returns GetBalanceOutputDto containing account ID and balance
   */
  execute(accountId: string): Promise<GetBalanceOutputDto>;
}