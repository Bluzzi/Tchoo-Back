### [BACK TO API](../../API.md)

**``POST`` /api/pets/get_top**  
Get an account's owned pets from the token.

#### Request
| Name        | Type | Require | Descritpion                                           |
| ----------- | ---- | ------- | ----------------------------------------------------- |
| start_index | int  | true    | The start index from which to retrieve the nfts       |
| stop_index  | int  | true    | The stop index from which to stop retrieving the nfts |

#### Response
##### Success
```json
{
  "success": true,
  "total_count": 1000,
  "top_nfts": [
      {
      "nonce": 0,
      "two_d_picture": "https://urltopicturepreview.com",
      "name": "Jerry",
      "holder_username": "greg",
      "prestige_balance": 10
    }
  ]
}
```