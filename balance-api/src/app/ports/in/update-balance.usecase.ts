/**
 * Interface for updating account balance use case
 */

import { UpdateBalanceInputDto } from "src/app/dto/input/update-balance-input.dto";

export interface IUpdateBalanceUseCase {

  /**
   * Executes the use case to update the balance of an account
   * @param input UpdateBalanceInputDto containing account ID and new balance
   */
  execute(input: UpdateBalanceInputDto): Promise<void>;
}