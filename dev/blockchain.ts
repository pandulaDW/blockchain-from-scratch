import { Block, Transaction } from "./models";

class Blockchain {
  chain: Block[];
  pendingTransactions: Transaction[];

  constructor() {
    this.chain = [];
    this.pendingTransactions = [];
  }

  createNewBlock(nonce: any, previousBlockHash: string, Hash: string): Block {
    const newBlock: Block = {
      index: this.chain.length + 1,
      timestamp: Date.now(),
      transactions: this.pendingTransactions,
      nonce,
      hash: Hash,
      previousBlockHash,
    };

    this.pendingTransactions = [];
    this.chain.push(newBlock);

    return newBlock;
  }

  getLastBlock() {
    return this.chain[this.chain.length - 1];
  }

  // returns the number of the block that this transaction will be added to
  createNewTransaction(amount: number, sender: string, recipient: string) {
    const newTransaction: Transaction = {
      amount,
      sender,
      recipient,
    };

    this.pendingTransactions.push(newTransaction);

    return this.getLastBlock()["index"] + 1;
  }
}

export default Blockchain;
