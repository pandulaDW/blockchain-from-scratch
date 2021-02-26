class Blockchain {
  chain: any[];
  newTransactions: any[];

  constructor() {
    this.chain = [];
    this.newTransactions = [];
  }

  createNewBlock(nonce: any, previousBlockHash: string, Hash: string) {
    const newBlock = {
      index: this.chain.length + 1,
      timestamp: Date.now(),
      transactions: this.newTransactions,
      nonce,
      hash: Hash,
      previousBlockHash,
    };

    this.newTransactions = [];
    this.chain.push(newBlock);

    return newBlock;
  }
}
