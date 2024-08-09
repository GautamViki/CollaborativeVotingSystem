package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"voting_system/dto"
	redisdriver "voting_system/redisDriver"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

type user struct {
}

func NewUserClient() *user {
	return &user{}
}

func (u *user) SetUser(w http.ResponseWriter, r *http.Request) {
	db := redisdriver.ConnectRedis()
	user := dto.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Got error while decoding body")
		return
	}
	ctx := r.Context()
	db.Set(ctx, "username", user.Username, 0)
	db.Set(ctx, "password", user.Password, 0)
}

func (u *user) CastVote(w http.ResponseWriter, r *http.Request) {
	// db := redisdriver.ConnectRedis()
	user := dto.User{}
	// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	// 	fmt.Println("Got error while decoding body")
	// 	return
	// }
	// ctx := r.Context()
	// fmt.Println("-------------------------", db.Get(ctx, "username"))

	decoded := context.Get(r, "decoded")
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}
