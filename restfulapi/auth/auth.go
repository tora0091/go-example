package auth

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"go-example/restfulapi/responsebody"
)

func GetAuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256) // jwt.SigningMethodHS256

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["name"] = "master"              // sammple name (get database)
	claims["email"] = "master@example.com" // samle email (get database)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// this-is-your-secret-key: set env .bashrc or database or login system
	// export SIGNINGKEY=this-is-your-secret-key, you get os.Getenv("SIGNINGKEY")
	tokenString, err := token.SignedString([]byte("this-is-your-secret-key"))
	if err != nil {
		log.Fatalln(err)
	}
	w.Write([]byte(tokenString))
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("this-is-your-secret-key")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			responsebody.StatusUnauthorized(w)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			responsebody.StatusUnauthorized(w)
			return
		}

		// check database or set session ...
		authorized := claims.(jwt.MapClaims)["authorized"].(bool)
		name := claims.(jwt.MapClaims)["name"].(string)
		email := claims.(jwt.MapClaims)["email"].(string)

		log.Printf("Authorized: %t\n", authorized)
		log.Printf("Name: %s\n", name)
		log.Printf("Email: %s\n", email)

		next.ServeHTTP(w, r)
	})
}
