# Documentation API

## Information
All requests are made with the "POST" method, the data must be sent in JSON as a string.  

The parameter "private_token" represents this token which is totally private :  
  - <details>
    <summary>Private secret token</summary>
    00ed35b450dc8a87cd7f22ee838c51e85617d6fe2bfae43c92be5884811b3600  
  </details>  
  
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
    - [GET_INFOS_BY_DISCORD](./routes/account/GET_INFOS_BY_DISCORD.md)
    - [GET_INFOS_BY_TOKEN](./routes/account/GET_INFOS_BY_TOKEN.md)
    - [GET_INFOS_BY_USERNAME](./routes/account/GET_INFOS_BY_USERNAME.md)
    - [SET_WHITELISTED](./routes/account/SET_WHITELISTED.md)
- ``/api/invites/``
    - [ADD_INVITE](./routes/invites/ADD_INVITE.md)
    - [GET_INVITES](./routes/invites/GET_INVITES.md)
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