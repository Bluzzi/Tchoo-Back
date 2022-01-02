### [BACK TO API](../API.md)

**``POST`` /api/account/link_wallet**  
The entry-point to link an elrond wallet with an account.

#### Request
| Name      | Type   | Require | Descritpion                                      |
| --------- | ------ | ------- | ------------------------------------------------ |
| token     | string | true    | The account token                                |
| signature | string | true    | The signature returned by the elrond login       |
| address   | string | true    | The address to which we need to link the account |

#### Response

###### Success
```json
{
  "success": true,
}
```

###### Error
```json
{
  "success": false,
  "error": "error id"
}
```

##### Error List
| Error Id              | Meaning                                    |
| --------------------- | ------------------------------------------ |
| account.token_invalid | Token is invalid                           |
| field.empty           | One of the fields is empty                 |
| wallet.used           | Wallet is already used in an other account |