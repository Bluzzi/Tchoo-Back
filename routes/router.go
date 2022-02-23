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
	mainRouter := r.PathPrefix("/api").Subrouter()

	accountRouter := mainRouter.PathPrefix("/account").Subrouter()
	accountRouter.HandleFunc("/create", account.HandleCreateAccount).Methods("POST")
	accountRouter.HandleFunc("/login", account.HandleLoginAccount).Methods("POST")
	accountRouter.HandleFunc("/logout", account.HandleLogoutAccount).Methods("POST")
	accountRouter.HandleFunc("/link_wallet", account.HandleLinkWallet).Methods("POST")
	accountRouter.HandleFunc("/is_wallet_linked", account.HandleIsWalletLinked).Methods("POST")
	accountRouter.HandleFunc("/get_infos", account.HandleGetInfos).Methods("POST")
	accountRouter.HandleFunc("/is_token_valid", account.HandleIsTokenValid).Methods("POST")


	petsRouter := mainRouter.PathPrefix("/pets").Subrouter()
	petsRouter.HandleFunc("/get_owned", pets.HandleGetOwnedRequest).Methods("POST")
	petsRouter.HandleFunc("/get", pets.HandleGetRequest).Methods("POST")
	petsRouter.HandleFunc("/get_top", pets.HandleGetTopRequest).Methods("POST")
	// Note: this endpoint is private and requires a private key, it adds an already minted to track it
	petsRouter.HandleFunc("/track_minted_nft", pets.HandleTrackMintedNftRequest).Methods("POST")
	petsRouter.HandleFunc("/get_stats", pets.HandleTrackMintedNftRequest).Methods("POST")

	interactionsRouter := petsRouter.PathPrefix("/interactions").Subrouter()
	interactionsRouter.HandleFunc("/feed", interactions.HandleFeedPet).Methods("POST")
	interactionsRouter.HandleFunc("/wash", interactions.HandleWashPet).Methods("POST")
	interactionsRouter.HandleFunc("/pet", interactions.HandleCaressPet).Methods("POST")
	interactionsRouter.HandleFunc("/sleep", interactions.HandleSleepPet).Methods("POST")


	lotteryRouter := mainRouter.PathPrefix("/lottery").Subrouter()
	lotteryRouter.HandleFunc("/get", lottery.HandleGetRequest).Methods("POST")
	lotteryRouter.HandleFunc("/buy_ticket", lottery.HandleBuyTicket).Methods("POST")

	fmt.Println("[LOAD] - Router loaded.")
	log.Fatal(http.ListenAndServe(":8080", cors.Default().Handler(r)))
}
