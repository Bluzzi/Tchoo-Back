# Documentation API

## Information
All of the ``GET`` parameters are passed through the URL.  
All of the ``POST`` data is passed through the body as JSON.  

## Routes
- ``/api/account/``
    - [LOGIN](./routes/account/LOGIN.md)
    - [LOGOUT](./routes/account/LOGOUT.md)
    - [CREATE](./routes/account/REGISTER.md)
    - [LINK_WALLET](./routes/account/LINK_WALLET.md)
    - [IS_WALLET_LINKED](./routes/account/IS_WALLET_LINKED.md)
    - [IS_TOKEN_VALID](./routes/account/IS_TOKEN_VALID.md)
    - [GET_INFOS](./routes/account/GET_INFOS.md)

- ``/api/pets/``
    - [GET](./routes/pets/GET.md)
    - [GET_OWNED](./routes/pets/GET_OWNED.md)
    - [GET_TOP](./routes/pets/GET_TOP.md)

- ``/api/interactions/``
    - [CARESS](./routes/pets/interactions/INTERACTION.md)
    - [FEED](./routes/pets/interactions/INTERACTION.md)
    - [SLEEP](./routes/pets/interactions/INTERACTION.md)
    - [WASH](./routes/pets/interactions/INTERACTION.md)
