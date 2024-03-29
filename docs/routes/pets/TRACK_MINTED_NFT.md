### [BACK TO API](../../API.md)

**``POST`` /api/pets/track_minted_nft**  
Track minted nft's private endpoint.

#### Request
| Name                         | Type   | Require | Descritpion                 |
| ---------------------------- | ------ | ------- | --------------------------- |
| private_token                | string | true    | The secret key              |
| nonce                        | int    | true    | Nft nonce                   |
| three_d_model                | string | true    | Url to 3d model             |
| two_d_picture                | string | true    | Url to 2d rendered image    |
| name                         | string | true    | Nft name                    |
| animations                   | Object | true    | Nft's animations            |
| points_balance               | int    | true    | Nft's lottery point balance |
| prestige_balance             | int    | true    | Nft's prestige balance      |
| points_per_five_minutes_base | int    | true    | Points won base             |
| points_per_five_minutes_real | int    | true    | Points won for real         |


#### Response
##### Success
```json
{
  "success": true
}
```