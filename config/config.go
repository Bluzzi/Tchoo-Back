package config

import (
	"encoding/json"
	"fmt"
	"github.com/ElrondNetwork/elrond-sdk/erdgo"
	"io/ioutil"
)

type TConfig struct {
	Password string `json:"password"`
	SecretJson string `json:"secret-json"`
	PointsForATicket float64 `json:"points-for-a-ticket"`
	CollectionToken string `json:"collection-token-identifier"`
	ContractAddress string `json:"contract-address"`
	PrivateKey []byte
}

var Config TConfig

func Load()  {
	file, _ := ioutil.ReadFile("./data/config.json")
	_ = json.Unmarshal(file, &Config)

	pk, _ := erdgo.LoadPrivateKeyFromJsonFile(Config.SecretJson, Config.Password)
	Config.PrivateKey = pk
	fmt.Println("[LOAD] - Configuration loaded.")
}
