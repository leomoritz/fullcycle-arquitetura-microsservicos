import { Controller, Get, HttpException, HttpStatus, NotFoundException, Param, Inject } from '@nestjs/common';
import * as usecase from 'src/app/ports/in/get-balance.usecase';
import { DomainException } from 'src/shared/exceptions/domain.exception';

/**
 * Balance Controller
 * Handles HTTP requests for balance-related operations
 */
@Controller('balances')
export class BalanceController {
  constructor(
    @Inject('IGetBalanceUseCase')
    private readonly getBalanceUseCase: usecase.IGetBalanceUseCase,
  ) {}

  /**
   * Get the balance of an account by ID
   * @param accountId ID of the account
   * @returns Balance information of the account
   * @throws HttpException if account not found or validation fails
   */
  @Get(':accountId')
  async getBalance(@Param('accountId') accountId: string) {
    try {
      return await this.getBalanceUseCase.execute(accountId);
    } catch (error: any) {
      if (error instanceof NotFoundException) {
        throw new HttpException(error.message, HttpStatus.NOT_FOUND);
      }
      if (error instanceof DomainException) {
        throw new HttpException(error.message, HttpStatus.BAD_REQUEST);
      }
      throw new HttpException('Unexpected error', HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }
}
