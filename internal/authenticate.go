package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"voting_system/dto"
	redisdriver "voting_system/redisDriver"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

var secretKey = []byte("secret-key")

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	db := redisdriver.ConnectRedis()
	ctx := r.Context()
	user := dto.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Got error while decoding body")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Username,
			"password": user.Password,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Got error while preparing token"))
	}
	expiresAt := time.Now().Add(time.Minute * 1).Unix()
	db.Set(ctx, "token", tokenString, 0)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}

func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r := http.Request{}
		ctx := r.Context()
		authorizationHeader := req.Header.Get("token")
		db := redisdriver.ConnectRedis()
		username := db.Get(ctx, "username").Val()
		token, error := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(secretKey), nil
		})
		if error != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Error in parsing token"))
			return
		}
		if token.Valid {
			var user dto.User
			mapstructure.Decode(token.Claims, &user)
			fmt.Println(username, user.Username)
			if username != user.Username {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte("Invalid authorization token - Does not match UserID"))
				return
			}

			context.Set(req, "decoded", token.Claims)
			next(w, req)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Invalid authorization token"))
		}
	})
}
