package routes

import (
	"MetaFriend/routes/account"
	"fmt"
	"github.com/gorilla/mux"
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

	fmt.Println("[LOAD] - Router loaded.")
	log.Fatal(http.ListenAndServe(":8080", r))
}
