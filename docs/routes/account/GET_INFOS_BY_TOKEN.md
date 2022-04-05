### [BACK TO API](../../API.md)

**``POST`` /api/account/get_infos_by_token**  
Get an account's information from the token.

#### Request
| Name  | Type   | Require | Descritpion                                                       |
| ----- | ------ | ------- | ----------------------------------------------------------------- |
| token | string | true    | The account's token that we need to retrieve the information from |

#### Response
##### Success
```json
{
  "username": "Swourire",
  "unique_username": "swourire",
  "wallet": "erd1aiejabiazebfjfkfuzuzenre1é19129ujpma",
  "is_whitelisted": true,
  "owned_pets": [],
  "discord_id": "191289812892198"
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
| Error Id              | Meaning                                    |
| --------------------- | ------------------------------------------ |
| account.token_invalid | Token is invalid                           |
| field.empty           | One of the fields is empty                 |