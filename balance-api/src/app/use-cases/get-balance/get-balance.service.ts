/** Service for getting account balance (implements GetBalanceUseCase) */

import { Injectable, NotFoundException, Inject } from "@nestjs/common";
import { IGetBalanceUseCase } from "../../ports/in/get-balance.usecase";
import { GetBalanceOutputDto } from "src/app/dto/output/get-balance-output.dto";
import * as port from "src/app/ports/out/accounts-balance.gateway";
import { AccountMapper } from "src/app/mappers/account.mapper";

@Injectable()
export class GetBalanceService implements IGetBalanceUseCase {
  constructor(
    @Inject('IAccountsBalanceGateway')
    private readonly accountsBalanceGateway: port.IAccountsBalanceGateway,
  ) { }

 
  async execute(accountId: string): Promise<GetBalanceOutputDto> {
    const balance = await this.accountsBalanceGateway.getBalanceByAccountId(accountId);
    if (balance == null || balance == undefined) {
      throw new NotFoundException("Não foi possível recuperar o saldo atual para conta informada")
    }
    return AccountMapper.toGetBalanceOutputDTO(accountId, balance);
  }
}