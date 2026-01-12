// parser.js
import { Command } from './command.js';

class PaymentParser {
  static parse(input) {
    const paymentCommand = new Command(input);
    return paymentCommand.execute();
  }
}

class Command {
  constructor(input) {
    this.input = input;
  }

  execute() {
    const { amount, currency } = this.input;
    if (!amount || !currency) {
      throw new Error('Invalid input: missing or invalid amount and/or currency');
    }

    const paymentAmount = parseFloat(amount);
    if (isNaN(paymentAmount)) {
      throw new Error('Invalid amount: must be a number');
    }

    return {
      amount: paymentAmount,
      currency,
    };
  }
}

export { PaymentParser };