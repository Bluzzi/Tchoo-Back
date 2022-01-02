### [BACK TO API](../API.md)

**``POST`` /api/account/login**  
The entry-point to retrieve a login token which can expire and can be used to interact with the rest of the API.

#### Request
| Name     | Type   | Require | Descritpion                                       |
| -------- | ------ | ------- | ------------------------------------------------- |
| username | string | true    | The login username                                |
| password | string | true    | The login password as text (hash is server sided) |

#### Response

###### Success
```json
{
  "success": true,
  "token": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx-expirationtimestamp"
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
| Error Id              | Meaning                                     |
| --------------------- | ------------------------------------------- |
| account.invalid_login | The supplied login informations are invalid |
| field.empty           | One of the fields is empty                  |

