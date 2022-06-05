/*
Package middlewares holds all the functions that are used for validating a token (JWT-based).
*/

package middlewares

import (
	"errors"
	"net/http"
	"os"
	"passer-auth-service-v4/pkg/utils"
	"strings"
)

// ValidateToken is a middleware that will check for the presence of a 'Token' attribute in the request header.
// It will permit the request to continue its flow to the secureed api endpoint if the 'Token' is present and valid.
// A valid 'Token' must satisfy the following:
// - the signature segment of the 'Token' must be consistent when this middleware signs the content of the Header and Payload segments (of the 'Token') with the secret key;
// - the 'exp' attribute in the Payload (encoded in Base64 format) has not come to pass.
func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// this configuration setting was first loaded from the .env file
		// in the main() function.
		JWT_SECRET_KEY := os.Getenv("JWT_SECRET_KEY")

		// get the jwt from the request header.
		authorization := req.Header.Get("Authorization")
		if strings.TrimSpace(authorization) == "" {
			res.Header().Set("Content-Type", "application/json")
			customErr := errors.New(`[MW]: no token found`)
			utils.SendForbiddenMsgToClient(&res, customErr)
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		if strings.TrimSpace(token) == "" {
			// empty token
			res.Header().Set("Content-Type", "application/json")
			customErr := errors.New(`[MW]: no token found`)
			utils.SendForbiddenMsgToClient(&res, customErr)
			return
		}

		// ok.
		// jwt validation logic here.
		ok, err := utils.Verify(token, JWT_SECRET_KEY)
		if err != nil {
			res.Header().Set("Content-Type", "application/json")
			customErr := errors.New(`[MW-JWT]: token is not valid`)
			utils.SendForbiddenMsgToClient(&res, customErr)
			return
			//
		}

		if !ok {
			res.Header().Set("Content-Type", "application/json")
			customErr := errors.New(`[MW-JWT] token is not valid`)
			utils.SendForbiddenMsgToClient(&res, customErr)
			return
			//
		}

		// ok. valid token. direct the request to the next handler.
		next(res, req)
	}
}
