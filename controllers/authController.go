package controllers
//
//import (
//	"encoding/json"
//	"errors"
//	"github.com/Muart-C/webcbo/models"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gorilla/context"
//	"log"
//	"net/http"
//	"strings"
//)
//
////GetAuthToken controller
//func GetAuthToken(w http.ResponseWriter, r *http.Request) {
//	var user models.User
//
//	_ = json.NewDecoder(r.Body).Decode(&user)
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"userName": user.UserName,
//		"password": user.Password,
//	})
//
//	tokenString, err := token.SignedString([]byte(".12tad3sa3dfe2349$##@^(>./"))
//	if err != nil {
//		log.Fatal(err)
//	}
//	json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})
//}
//
////AuthenticationMiddleware controller
//func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		authHeaders := r.Header.Get("Authorization")
//		if authHeaders != " " {
//			bearerToken := strings.Split(authHeaders, " ")
//			if len(bearerToken) == 2 {
//				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
//					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//						return nil, errors.New("n error occurred while parsing the token")
//					}
//					return []byte(".12tad3sa3dfe2349$##@^(>./"), nil
//				})
//				if err != nil {
//					RespondWithError(w, http.StatusHTTPVersionNotSupported, "an error occurred")
//				}
//				if token.Valid {
//					log.Println("Token was valid")
//					context.Set(r, "decoded", token.Claims)
//					next(w, r)
//				} else {
//					RespondWithError(w, http.StatusBadRequest, "invalid request token")
//				}
//			}
//		} else {
//			log.Println("auth token is required")
//			RespondWithError(w, http.StatusNotFound, "an authorization header is required")
//		}
//	})
//}
