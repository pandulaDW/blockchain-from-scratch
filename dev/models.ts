export interface Block {
  index: number;
  timestamp: number;
  transactions: any[];
  nonce: any;
  hash: string;
  previousBlockHash: string;
}
