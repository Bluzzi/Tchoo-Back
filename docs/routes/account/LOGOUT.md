### [BACK TO API](../../API.md)

**``POST`` /api/account/logout**  
The entry-point to retrieve a login token which can expire and can be used to interact with the rest of the API.

#### Request
| Name  | Type   | Require | Descritpion             |
| ----- | ------ | ------- | ----------------------- |
| token | string | true    | The token to invalidate |

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
| Error Id              | Meaning                       |
| --------------------- | ----------------------------- |
| account.token_invalid | The supplied token is invalid |
| field.empty           | One of the fields is empty    |

