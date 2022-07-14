import { TokenConfiguration } from "@cosmwasm/bcp";

export const binaryName = "cosmwasm-faucet";
export const concurrency: number = Number.parseInt(process.env.FAUCET_CONCURRENCY || "", 10) || 5;
export const port: number = Number.parseInt(process.env.FAUCET_PORT || "", 10) || 8000;
export const mnemonic: string | undefined = process.env.FAUCET_MNEMONIC;

export const addressPrefix = "secret";

/** For the local development chain */
export const developmentTokenConfig: TokenConfiguration = {
  bankTokens: [
    {
      fractionalDigits: 6,
      name: "Fee Token",
      ticker: "SCRT",
      denom: "uscrt",
    },
  ],
};
