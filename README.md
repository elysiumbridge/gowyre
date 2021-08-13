# gowyre

Golang unofficial Wyre API client. Check out `https://docs.sendwyre.com/reference` for full API specification.

## Example

```go

```

## Supported

### Key Management

- [] POST `https://api.sendwyre.com/v2/sessions/auth/key`
- [] POST `https://api.sendwyre.com/v2/apiKeys`
- [] DELETE `https://api.sendwyre.com/v2/apiKey/apiKey`

### Wyre Checkout

- [x] POST `https://api.sendwyre.com/v3/widget/limits/calculate`
- [x] POST `https://api.sendwyre.com/v3/orders/reserve`
- [x] POST `https://api.sendwyre.com/v3/orders/quote/partner`
- [x] GET `https://api.sendwyre.com/v3/orders/orderId`
- [x] GET `https://api.sendwyre.com/v3/orders/id/full`
- [x] GET `https://api.sendwyre.com/v3/orders/list`
- [x] GET `https://api.sendwyre.com/v2/transfer/transferId/track`
- [x] GET `https://api.sendwyre.com/v3/widget/supportedCountries`
- [x] GET `https://api.sendwyre.com/v3/orders/reservation/reservationId`

### Card Processing

- [x] POST `https://api.sendwyre.com/v3/debitcard/process/partner`
- [x] GET `https://api.sendwyre.com/v3/debitcard/authorization/orderId`
- [x] POST `https://api.sendwyre.com/v3/debitcard/authorize/partner`
- [x] POST `https://api.sendwyre.com/v3/orders/:orderId/refund/partner`
- [] POST `https://api.sendwyre.com/v3/apple-pay/process/partner`

### Wallets

- [x] POST `https://api.sendwyre.com/v2/wallets`
- [x] POST `https://api.sendwyre.com/v2/wallets/batch`
- [x] GET `https://api.sendwyre.com/v2/wallet/:walletId`
- [x] POST `https://api.sendwyre.com/v2/wallet/:walletId/update`
- [x] DELETE `https://api.sendwyre.com/v2/wallet/:walletId`
- [x] GET `https://api.sendwyre.com/v2/wallets`

### Savings

- [] POST `https://example.com/v2/wallets`
- [] GET `https://example.com/v2/wallet/walletId`

### Transfers and Exchanges

- [x] POST `https://api.sendwyre.com/v3/transfers`
- [x] POST `https://api.sendwyre.com/v3/transfers/transferId:/confirm`
- [x] GET `https://api.sendwyre.com/v3/transfers/transferId`
- [x] GET `https://api.sendwyre.com/v2/transfer?customId=XXXXXXX`
- [x] GET `https://api.sendwyre.com/v3/transfers`
- [x] GET `https://api.sendwyre.com/v3/rates`
- [x] POST `https://api.sendwyre.com/v3/swaps`
- [x] GET `https://api.sendwyre.com/v3/swaps/id`
- [x] POST `https://api.sendwyre.com/v3/swaps/ipCheck`

### Accounts

- [] POST `https://api.sendwyre.com/v3/accounts`
- [] GET `https://api.sendwyre.com/v3/accounts/accountId`
- [] GET `https://api.sendwyre.com/v2/account`
- [] POST `https://api.sendwyre.com/v3/accounts/accountId`
- [] POST `https://api.sendwyre.com/v3/accounts/accountId/fieldId`
- [] GET `https://api.sendwyre.com/v3/limits`
- [] GET `https://api.sendwyre.com/v3/accounts/accountId/statusHistory`
- [] GET `https://api.sendwyre.com/v3/accounts/accountId/profileFieldsStatuses`
- [] GET `https://api.sendwyre.com/v3/accounts/accountId/referredAccounts`

### Payment Methods

- [] POST `https://example.com/v2/paymentMethods`
- [] GET `https://api.sendwyre.com/v2/transfer/transferId`
- [] POST `https://example.com/v2/paymentMethods`
- [] POST `https://api.sendwyre.com/v2/paymentMethod/paymentMethodId/followup`
- [] DELETE `https://api.sendwyre.com/v2/paymentMethod/paymentMethodId/followup`
- [] GET `https://api.sendwyre.com/v2/paymentMethods`
- [] GET `https://api.sendwyre.com/v2/paymentMethod/paymentMethodId`
- [] POST `https://api.sendwyre.com/v2/paymentMethod/:paymentMethodId/attach`

### Global Payouts

- [] POST `https://api.sendwyre.com/v2/transfers`
- [] GET `https://api.sendwyre.com/v:version/bankinfo/country`
