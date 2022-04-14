### [BACK TO API](../../API.md)

**``POST`` /api/account/get_infos_by_discord**  
Get an account's information from the discord id.

#### Request
| Name       | Type   | Require | Descritpion                                                    |
| ---------- | ------ | ------- | -------------------------------------------------------------- |
| discord_id | string | true    | The account's id that we need to retrieve the information from |

#### Response
##### Success
```json
{
  "username": "Swourire",
  "unique_username": "swourire",
  "wallet": "erd1aiejabiazebfjfkfuzuzenre1é19129ujpma",
  "is_whitelisted": true,
  "owned_pets": [0, 12, 13]
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