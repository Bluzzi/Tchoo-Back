### [BACK TO API](../../API.md)

**``POST`` /api/account/get_infos**  
Get an account's information from the token

#### Request
| Name  | Type   | Require | Descritpion                                                       |
| ----- | ------ | ------- | ----------------------------------------------------------------- |
| token | string | true    | The account's token that we need to retrieve the information from |


#### Response

###### Success
```json
{
  "username": "Swourire",
  "unique_username": "swourire",
  "wallet": "erd1aiejabiazebfjfkfuzuzenre1é19129ujpma"
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