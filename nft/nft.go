package nft

import (
	"MetaFriend/config"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ApiUrl = "https://devnet-api.elrond.com/"
)

type Nft struct {
	Nonce int64 `json:"nonce"`
	Owner string `json:"owner"`
}

func UpdateNonceAddressCache()  {
	fmt.Println("[LOAD] - Cache loaded.")

	var nfts []Nft
	c := &http.Client{}
	resp, e := c.Get(ApiUrl + "nfts?collection=" + config.Config.CollectionToken + "&type=NonFungibleESDT&withOwner=true")

	if e != nil {
		fmt.Println(e)
		return
	}

	e = json.NewDecoder(resp.Body).Decode(&nfts)
	if e != nil {
		fmt.Println(e)
		return
	}

	_ddressNoncesCache := map[string][]int64{}
	for _, nft := range nfts {
		NonceAddressCache[nft.Nonce] = nft.Owner
		if _, exists := _ddressNoncesCache[nft.Owner]; exists {
			_ddressNoncesCache[nft.Owner] = append(_ddressNoncesCache[nft.Owner], nft.Nonce)
		} else {
			_ddressNoncesCache[nft.Owner] = []int64{ nft.Nonce }
		}
	}

	AddressNoncesCache = _ddressNoncesCache
}