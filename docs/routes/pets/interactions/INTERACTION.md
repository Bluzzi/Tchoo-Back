### [BACK TO API](../../API.md)

**``POST`` /api/pets/interactions/pet**  
Caress the NFT interaction

**``POST`` /api/pets/interactions/feed**  
Feed the NFT interaction

**``POST`` /api/pets/interactions/pet**  
Caress the NFT interaction

**``POST`` /api/pets/interactions/sleep**  
Sleep the NFT interaction

#### Request
| Name      | Type   | Require | Descritpion         |
| --------- | ------ | ------- | ------------------- |
| token     | string | true    | The account's token |
| pet_nonce | int    | true    | The pet's nonce     |


#### Response

###### Success
```json
{
  "success": true,
  "error": ""
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
| Error Id              | Meaning                          |
| --------------------- | -------------------------------- |
| account.token_invalid | Token is invalid                 |
| pet.not_owned         | Pet is not owned by the guy      |
| field.empty           | One of the fields is empty       |
| pet.action_cooldown:  | The pet's actions is on cooldown |