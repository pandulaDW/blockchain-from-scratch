export interface Block {
  index: number;
  timestamp: number;
  transactions: any[];
  nonce: any;
  hash: string;
  previousBlockHash: string;
}

export interface Transaction {
  amount: number;
  sender: string;
  recipient: string;
}
