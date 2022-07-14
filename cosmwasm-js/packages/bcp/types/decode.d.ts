import { IndexedTx, types, Coin } from "@cosmwasm/sdk";
import {
  Amount,
  ChainId,
  ConfirmedAndSignedTransaction,
  ConfirmedTransaction,
  Fee,
  FullSignature,
  Nonce,
  PubkeyBundle,
  SignatureBytes,
  SignedTransaction,
  UnsignedTransaction,
} from "@iov/bcp";
import { Decimal } from "@iov/encoding";
import { BankTokens, Erc20Token } from "./types";
export declare function decodePubkey(pubkey: types.PubKey): PubkeyBundle;
export declare function decodeSignature(signature: string): SignatureBytes;
export declare function decodeFullSignature(signature: types.StdSignature, nonce: number): FullSignature;
export declare function coinToDecimal(tokens: BankTokens, coin: Coin): readonly [Decimal, string];
export declare function decodeAmount(tokens: BankTokens, coin: Coin): Amount;
export declare function parseMsg(
  msg: types.Msg,
  memo: string | undefined,
  chainId: ChainId,
  tokens: BankTokens,
  erc20Tokens: readonly Erc20Token[],
): UnsignedTransaction;
export declare function parseFee(fee: types.StdFee, tokens: BankTokens): Fee;
export declare function parseUnsignedTx(
  txValue: types.StdTx,
  chainId: ChainId,
  tokens: BankTokens,
  erc20Tokens: readonly Erc20Token[],
): UnsignedTransaction;
export declare function parseSignedTx(
  txValue: types.StdTx,
  chainId: ChainId,
  nonce: Nonce,
  tokens: BankTokens,
  erc20Tokens: readonly Erc20Token[],
): SignedTransaction;
export declare function parseTxsResponseUnsigned(
  chainId: ChainId,
  currentHeight: number,
  response: IndexedTx,
  tokens: BankTokens,
  erc20Tokens: readonly Erc20Token[],
): ConfirmedTransaction<UnsignedTransaction>;
export declare function parseTxsResponseSigned(
  chainId: ChainId,
  currentHeight: number,
  nonce: Nonce,
  response: IndexedTx,
  tokens: BankTokens,
  erc20Tokens: readonly Erc20Token[],
): ConfirmedAndSignedTransaction<UnsignedTransaction>;
