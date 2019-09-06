package routes

import (
	"encoding/json"
	"eth-rbac/auth"
	"eth-rbac/models"
	"fmt"
	"net/http"
	"strings"
)

const SIGTXT = "Sign into as %v to Frontier backend services. By signing that you accept our terms and conditions and you understand that Frontier and all the related entities have no responsibilities on your private key, your eth, or your tokens."

type EthRequest struct {
	Account string `json:"account"`
	Signed  string `json:"signed"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var esr EthRequest
	err := decoder.Decode(&esr)
	if err != nil {
		w.Write([]byte("Error while decoding request body"))
		return
	}
	msg := fmt.Sprintf(SIGTXT, strings.ToLower(esr.Account))
	if !auth.VerifySig(esr.Account, esr.Signed, []byte(msg)) {
		w.Write([]byte("Hey! That's an invalid signature. Maybe signed the wrong message?"))
		return
	}
	u, err := models.GetUserByAddress(strings.ToLower(esr.Account))
	if err != nil {
		w.Write([]byte("User not found"))
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		panic("User couldn't be encoded, checkout the user model format.")
	}
	w.Write([]byte(uj))
}
