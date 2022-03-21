### [BACK TO API](../../API.md)

**``POST`` /api/pets/get_owned**  
Get an account's owned pets from the token.

#### Request
| Name               | Type   | Require | Descritpion                                                  |
| ------------------ | ------ | ------- | ------------------------------------------------------------ |
| is_wallet_supplied | bool   | true    | Do we use the wallet or the token to retrieve nfts?          |
| token              | string | true    | The account's token that we need to retrieve the owned nfts  |
| wallet             | string | true    | The account's wallet that we need to retrieve the owned nfts |

#### Response
##### Success
```json
{
  "success": true,
  "owned_nfts_nonces": [
      0,
      1,
      2,
      3
  ]
}
```

###### Error
```json
{
  "success": false,
  "error": "error id"
}
```

#### Error List
| Error Id                 | Meaning                         |
| ------------------------ | ------------------------------- |
| account.token_invalid    | Token is invalid                |
| account.no_wallet_linked | Account did not link the wallet |
| field.empty              | One of the fields is empty      |