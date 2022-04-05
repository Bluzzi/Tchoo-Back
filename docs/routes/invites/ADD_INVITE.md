> [DELETE ME] Don't forget to change the link below according to the number of sub-folders
### [BACK TO API](./API.md)

**``POST`` /api/invites/add_invite**  
Add an invite to a discord id

#### Request
| Name        | Type   | Require | Descritpion                |
| ----------- | ------ | ------- | -------------------------- |
| discord_id  | string | true    | L'id du discord à ajouter  |
| amount      | int    | true    | Combien d'invite à ajouter |
| private_key | string | true    | La clé privée de l'api     |

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
| Error Id    | Meaning                    |
| ----------- | -------------------------- |
| field.empty | One of the fields is empty |