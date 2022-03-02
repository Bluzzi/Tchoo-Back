### [BACK TO API](../../API.md)

**``POST`` /api/pets/get**  
Get a NFT's informations from it's nonce (id)

#### Request
| Name  | Type | Require | Descritpion          |
| ----- | ---- | ------- | -------------------- |
| nonce | int  | true    | The nft's nonce (id) |


#### Response

###### Success
```json
{
  "nonce": 0,
  "three_d_model": "https://urltomodel.com",
  "two_d_picture": "https://urltopicturepreview.com",
  "name": "Jerry",
  "holder_username": "greg",
  "points_balance": 1000,
  "prestige_balance": 10,
  "animations": {
    "idle": "https://url",
    "purring": "https://url"
  },
  //A constant
  "points_per_five_minutes_base": 10,
  //affected by the state of the pet, will be the passive revenue
  "points_per_five_minutes_real": 5,
  // Timestamps until when an action can be used
  "actions_used": {
      "wash": 1000,
      "feed": 1000,
      "pet": 1000,
      "sleep": 1000
  }
}
```