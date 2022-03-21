package routes

import (
	"MetaFriend/routes/account"
	"MetaFriend/routes/lottery"
	"MetaFriend/routes/pets"
	"MetaFriend/routes/pets/interactions"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Load()  {
	r := mux.NewRouter()
	mainRouter := r.PathPrefix("/v1").Subrouter()

	// Account :
	accountRouter := mainRouter.PathPrefix("/account").Subrouter()
	accountRouter.HandleFunc("/create", account.HandleCreateAccount).Methods("POST")
	accountRouter.HandleFunc("/login", account.HandleLoginAccount).Methods("POST")
	accountRouter.HandleFunc("/logout", account.HandleLogoutAccount).Methods("POST")
	accountRouter.HandleFunc("/link_wallet", account.HandleLinkWallet).Methods("POST")
	accountRouter.HandleFunc("/is_wallet_linked", account.HandleIsWalletLinked).Methods("POST")
	accountRouter.HandleFunc("/get_infos", account.HandleGetInfos).Methods("POST")
	accountRouter.HandleFunc("/is_token_valid", account.HandleIsTokenValid).Methods("POST")
	accountRouter.HandleFunc("/change_password", account.HandleChangePassword).Methods("POST")

	// Pets :
	petsRouter := mainRouter.PathPrefix("/pets").Subrouter()
	petsRouter.HandleFunc("/get_owned", pets.HandleGetOwnedRequest).Methods("POST")
	petsRouter.HandleFunc("/get", pets.HandleGetRequest).Methods("POST")
	petsRouter.HandleFunc("/get_top", pets.HandleGetTopRequest).Methods("POST")
	petsRouter.HandleFunc("/get_account_stats", pets.HandleGetAccountStats).Methods("POST")
	petsRouter.HandleFunc("/track_minted_nft", pets.HandleTrackMintedNftRequest).Methods("POST") // use the private key
	petsRouter.HandleFunc("/get_stats", pets.HandleTrackMintedNftRequest).Methods("POST")

	// Pets/Interactions :
	interactionsRouter := petsRouter.PathPrefix("/interactions").Subrouter()
	interactionsRouter.HandleFunc("/feed", interactions.HandleFeedPet).Methods("POST")
	interactionsRouter.HandleFunc("/wash", interactions.HandleWashPet).Methods("POST")
	interactionsRouter.HandleFunc("/pet", interactions.HandleCaressPet).Methods("POST")
	interactionsRouter.HandleFunc("/sleep", interactions.HandleSleepPet).Methods("POST")

	// Lottery :
	lotteryRouter := mainRouter.PathPrefix("/lottery").Subrouter()
	lotteryRouter.HandleFunc("/get", lottery.HandleGetRequest).Methods("POST")
	lotteryRouter.HandleFunc("/buy_ticket", lottery.HandleBuyTicket).Methods("POST")

	fmt.Println("[LOAD] - Router loaded.")
	log.Fatal(http.ListenAndServe(":3001", cors.Default().Handler(r)))
}
