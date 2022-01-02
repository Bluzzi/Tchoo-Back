### [BACK TO API](../../API.md)

**``POST`` /api/account/create**  
The entry-point to create an account.

#### Request
| Name     | Type   | Require | Descritpion                                         |
| -------- | ------ | ------- | --------------------------------------------------- |
| username | string | true    | The account username                                |
| password | string | true    | The account password as text (hash is server sided) |

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
| Error Id       | Meaning                                      |
| -------------- | -------------------------------------------- |
| account.exists | An account with this username already exists |
| field.empty    | One of the fields is empty                   |

