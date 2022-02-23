### [BACK TO API](../../API.md)

**``POST`` /api/account/is_wallet_linked**  
Know, with a token if a wallet is linked to account

#### Request
| Name  | Type   | Require | Description                                                              |
| ----- | ------ | ------- | ------------------------------------------------------------------------ |
| token | string | true    | The token linked to the account which we should check if it has a wallet |

#### Response

###### Success
```json
{
  "success": true || false,
  "error": ""
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
| Error Id              | Meaning                         |
| --------------------- | ------------------------------- |
| account.token_invalid | The supplied token doesn't work |
| field.empty           | One of the fields is empty      |

