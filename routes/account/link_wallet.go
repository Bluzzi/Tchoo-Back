package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"github.com/ElrondNetwork/elrond-go-crypto/signing"
	"github.com/ElrondNetwork/elrond-go-crypto/signing/ed25519"
	"github.com/ElrondNetwork/elrond-go-crypto/signing/ed25519/singlesig"
	"github.com/btcsuite/btcutil/bech32"
	"log"
	"net/http"
)

type LinkWalletRequest struct {
	Token string `json:"token"`
	Signature string `json:"signature"`
	Address string `json:"address"`
}

func (lR LinkWalletRequest) Verify() (bool, string) {
	if len(lR.Token) == 0 || len(lR.Signature) == 0 || len(lR.Address) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(lR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	if isWalletUsed, _ := authentication.Exists(authentication.FieldWallet, lR.Address); isWalletUsed {
		return false, errors.ErrorWalletUsed
	}

	return true, ""
}

func HandleLinkWallet(w http.ResponseWriter, r *http.Request)  {
	var linkWalletRequest LinkWalletRequest
	_ = json.NewDecoder(r.Body).Decode(&linkWalletRequest)

	if success, errorStr := linkWalletRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	suite := ed25519.NewEd25519()
	keyGenerator := signing.NewKeyGenerator(suite)
	_, bech32Decoded, _ := bech32.Decode(linkWalletRequest.Address)
	publicKey, err := keyGenerator.PublicKeyFromByteArray(bech32Decoded)
	if err != nil {
		return 
	}

	signer := &singlesig.Ed25519Signer{}
	errVerif := signer.Verify(publicKey, []byte(linkWalletRequest.Address + linkWalletRequest.Token + "{}"), []byte(linkWalletRequest.Signature))
	if errVerif != nil {
		log.Fatal(errVerif)
	}

	json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}

