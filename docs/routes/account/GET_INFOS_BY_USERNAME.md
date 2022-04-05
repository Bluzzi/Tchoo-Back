### [BACK TO API](../../API.md)

**``POST`` /api/account/get_infos_by_username**  
Get an account's information from the discord.

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
  "owned_pets": []
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
| Error Id             | Meaning                    |
| -------------------- | -------------------------- |
| account.not_existing | The account does not exist |
| field.empty          | One of the fields is empty |