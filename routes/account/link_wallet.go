package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ElrondNetwork/elrond-go-crypto/signing"
	"github.com/ElrondNetwork/elrond-go-crypto/signing/ed25519"
	"github.com/ElrondNetwork/elrond-go-crypto/signing/ed25519/singlesig"
	"github.com/ElrondNetwork/elrond-go/core/pubkeyConverter"
	"golang.org/x/crypto/sha3"
	"net/http"
	"strconv"
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
	converter, err := pubkeyConverter.NewBech32PubkeyConverter(32)
	if err != nil {
		fmt.Println(err)
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorWalletIncorrectSignature,
		})
		return
	}

	byteArray, err := converter.Decode(linkWalletRequest.Address)
	if err != nil {
		fmt.Println(err)
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorWalletIncorrectSignature,
		})
		return
	}

	publicKey, err := keyGenerator.PublicKeyFromByteArray(byteArray)
	if err != nil {
		fmt.Println(err)
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorWalletIncorrectSignature,
		})
		return
	}

	signer := &singlesig.Ed25519Signer{}


	message := linkWalletRequest.Address + linkWalletRequest.Token + "{}"

	var unsignedMessage []byte
	unsignedMessage = append(unsignedMessage, []byte("\x17Elrond Signed Message:\n")...)
	unsignedMessage = append(unsignedMessage, []byte(strconv.Itoa(len(message)))...)
	unsignedMessage = append(unsignedMessage, []byte(message)...)

	h := sha3.NewLegacyKeccak256()
	h.Write(unsignedMessage)

	unsignedMessage = h.Sum(nil)

	sigSrc := []byte(linkWalletRequest.Signature)
	sigDst := make([]byte, hex.DecodedLen(len(sigSrc)))
	_, _ = hex.Decode(sigDst, sigSrc)

	errVerif := signer.Verify(publicKey, unsignedMessage, sigDst)

	if errVerif != nil {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorWalletIncorrectSignature,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}