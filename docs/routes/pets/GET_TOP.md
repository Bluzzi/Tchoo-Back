### [BACK TO API](../../API.md)

**``POST`` /api/pets/get_top**  
Get an account's owned pets from the token

#### Request
| Name        | Type | Require | Descritpion                                           |
| ----------- | ---- | ------- | ----------------------------------------------------- |
| start_index | int  | true    | The start index from which to retrieve the nfts       |
| stop_index  | int  | true    | The stop index from which to stop retrieving the nfts |


#### Response
```json
{
  "success": true,
  "total_count": 1000,
  "top_nfts": [
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
  ]
}
```