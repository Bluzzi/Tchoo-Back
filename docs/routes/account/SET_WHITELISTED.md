### [BACK TO API](../../API.md)

**``POST`` /api/account/set_whitelisted**  
Set a person's whitelist status

#### Request
| Name        | Type   | Require | Description                                          |
| ----------- | ------ | ------- | ---------------------------------------------------- |
| discord_id  | string | true    | The discord id to which we need to set the whitelist |
| whitelist   | bool   | true    | Set or unset a whitelist                             |
| private_key | string | true    | The api's private key                                |

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
| Error Id             | Meaning                    |
| -------------------- | -------------------------- |
| account.not_existing | The account does not exist |
| field.empty          | One of the fields is empty |

