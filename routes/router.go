package routes

import (
	"encoding/json"
	"eth-rbac/auth"
	"eth-rbac/models"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

const SALT_TEXT = "Sign into as %v to Frontier backend services. By signing that you accept our terms and conditions and you understand that Frontier and all the related entities have no responsibilities on your private key, your eth, or your tokens."

type EthRequest struct {
	Account string `json:"account"`
	Signed  string `json:"signed"`
}

type Response struct {
	Result string `json:"result"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}

type Claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var esr EthRequest
	err := decoder.Decode(&esr)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	msg := fmt.Sprintf(SALT_TEXT, strings.ToLower(esr.Account))
	if !auth.VerifySig(esr.Account, esr.Signed, []byte(msg)) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	u, err := models.GetUserByAddress(strings.ToLower(esr.Account))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := Response{
		Result: "Connected",
		Role:   u.Role,
		Token:  tokenString,
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resj, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(resj))
}
