# Documentation API

## Information
All requests are made with the "POST" method, the data must be sent in JSON as a string.  

All requests responses will have at least its properties :
### Success
```json
{
  "success": true
}
```

### Error
```json
{
  "success": false,
  "error": "error id"
}
```

## Routes
- ``/api/account/``
    - [LOGIN](./routes/account/LOGIN.md)
    - [LOGOUT](./routes/account/LOGOUT.md)
    - [CREATE](./routes/account/REGISTER.md)
    - [LINK_WALLET](./routes/account/LINK_WALLET.md)
    - [IS_WALLET_LINKED](./routes/account/IS_WALLET_LINKED.md)
    - [IS_TOKEN_VALID](./routes/account/IS_TOKEN_VALID.md)
    - [CHANGE_PASSWORD](./routes/account/CHANGE_PASSWORD.md)
    - [GET_INFOS](./routes/account/GET_INFOS.md)
- ``/api/pets/``
    - [GET](./routes/pets/GET.md)
    - [GET_OWNED](./routes/pets/GET_OWNED.md)
    - [GET_TOP](./routes/pets/GET_TOP.md)
    - [GET_ACCOUNT_STATS](./routes/pets/GET_ACCOUNT_STATS.md)
    - [TRACK_MINTED_NFT](./routes/pets/TRACK_MINTED_NFT.md)
    - ``/interactions/``
        - [CARESS](./routes/pets/interactions/INTERACTIONS.md)
        - [FEED](./routes/pets/interactions/INTERACTIONS.md)
        - [SLEEP](./routes/pets/interactions/INTERACTIONS.md)
        - [WASH](./routes/pets/interactions/INTERACTIONS.md)