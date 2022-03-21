> [DELETE ME] Don't forget to change the link below according to the number of sub-folders
### [BACK TO API](./API.md)

**``POST`` /api/account/get_infos**  
Description of the request.

#### Request
| Name | Type | Require | Descritpion |
| ---- | ---- | ------- | ----------- |
|      |      |         |             |

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
| Error Id | Meaning                                                       |
| -------- | ------------------------------------------------------------- |
| error.id | Description of the error (it will be used on the client side) |