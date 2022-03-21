### [BACK TO API](../../API.md)

**``POST`` /api/pets/get_account_stats**  
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
  "username": "greg",
  "wallet": "erd13zdl72sfr8249y70c3ughhd3mpeay8d7e8m0nw849wzxeqq2hu8smwnsd5",
  "owned_nfts_nonces": [
    {
      "nonce": 0,
      "three_d_model": "https://urltomodel.com",
      "two_d_picture": "https://urltopicturepreview.com",
      "name": "Jerry",
      "holder_username": "greg",
      "points_balance": 1000,
      "prestige_balance": 10,
      "animations": {
        "idle": "https://url",
        "purring": "https://url"
      },
      "points_per_five_minutes_base": 10,
      "points_per_five_minutes_real": 5,
      "actions_used": {
        "wash": 1000,
        "feed": 1000,
        "pet": 1000,
        "sleep": 1000
      }
    }
  ]
}
```

##### Error
```json
{
  "success": false,
  "error": "error id"
}
```

#### Error List
| Error Id                 | Meaning                         |
| ------------------------ | ------------------------------- |
| account.username_invalid | Account username does not exist |
| field.empty              | One of the fields is empty      |