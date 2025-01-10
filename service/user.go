package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/upbreak/go-todo-app/auth"
	"github.com/upbreak/go-todo-app/store"
)

type GetUser struct {
	DB   store.Queryer
	Repo store.GetUserValidStore
	Jwt  *auth.JWTUtils
}

func (g *GetUser) GetUserValid(ctx context.Context, id string, pw string) (auth.JWTClaims, error) {
	hash := md5.Sum([]byte(pw))
	fmt.Println("hash: %v", hash)
	pwMd5 := hex.EncodeToString(hash[:])
	fmt.Println("pwMd5: %v", pwMd5)

	// 유저 확인
	user, err := g.Repo.GetUserValid(ctx, g.DB, id, pwMd5)
	if err != nil {
		return auth.JWTClaims{}, fmt.Errorf("fail to getUser: %w", err)
	}

	jwtClaims := auth.JWTClaims{
		UserId: user.UserId,
		Role:   auth.Admin,
	}

	// 토큰 생성
	tokenString, err := g.Jwt.GenerateToken(jwtClaims.UserId)
	if err != nil {
		return auth.JWTClaims{}, fmt.Errorf("fail to generate token: %w", err)
	}

	jwtClaims.Token = tokenString

	return jwtClaims, nil

}
