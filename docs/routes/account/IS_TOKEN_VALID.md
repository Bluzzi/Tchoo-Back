### [BACK TO API](../../API.md)

**``POST`` /api/account/is_token_valid**  
Is an account's token valid.

#### Request
| Name  | Type   | Require | Descritpion                                    |
| ----- | ------ | ------- | ---------------------------------------------- |
| token | string | true    | The token which we should check if it is valid |

#### Response
##### Success
```json
{
  "success": true,
}
```

##### Error
```json
{
  "success": false,
  "error": "error id"
}
```