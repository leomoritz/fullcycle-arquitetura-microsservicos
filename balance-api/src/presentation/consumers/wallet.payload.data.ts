/**
 * Class to define payload data structure
 */
export class WalletUpdateBalanceEventData {
    account_id_from: string;
    account_id_to: string;
    balance_account_id_from: number;
    balance_account_id_to: number;
}