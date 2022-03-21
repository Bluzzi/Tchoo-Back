### [BACK TO API](../../API.md)

**``POST`` /api/account/change_password**  
Change the account's password.

#### Request
| Name         | Type   | Require | Descritpion                                    |
| ------------ | ------ | ------- | ---------------------------------------------- |
| token        | string | true    | The token which we should check if it is valid |
| old_password | string | true    | Old password                                   |
| new_password | string | true    | New password                                   |

#### Response
##### Success
```json
{
  "success": true
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
| Error Id                   | Meaning                         |
| -------------------------- | ------------------------------- |
| account.token_invalid      | The supplied token doesn't work |
| account.password_incorrect | The old password doesn't match  |
| field.empty                | One of the fields is empty      |
