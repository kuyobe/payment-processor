"""
Payment Processor
================

A Python package for processing payments.

Installation
------------

To install the payment-processor package, run the following command:

```bash
pip install payment-processor
```

Usage
-----

To use the payment-processor package, import it in your Python script and create an instance of the PaymentProcessor class.

```python
from payment_processor import PaymentProcessor

processor = PaymentProcessor()
processor.process_payment("1234567890", 100.00)
```

API Documentation
----------------

### PaymentProcessor

*   `process_payment(card_number: str, amount: float)`: Process a payment using the provided card number and amount.

### Exceptions

*   `InvalidCardNumberError`: Raised when the provided card number is invalid.
*   `InsufficientFundsError`: Raised when the account has insufficient funds.

"""