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

func (g *GetUser) GetUserValid(ctx context.Context, id string, pw string) (*auth.JWTClaims, error) {
	// 비밀번호 암호화. postman을 이용한 테스트 때문에 여기서 암호화 함.
	hash := md5.Sum([]byte(pw))
	pwMd5 := hex.EncodeToString(hash[:])

	// 유저 확인
	user, err := g.Repo.GetUserValid(ctx, g.DB, id, pwMd5)
	if err != nil {
		return &auth.JWTClaims{}, fmt.Errorf("fail to getUser: %w", err)
	}

	jwtClaims := &auth.JWTClaims{
		UserId: user.UserId,
		Role:   auth.Admin,
	}

	// 토큰 생성
	jwtClaims, err = g.Jwt.GenerateToken(jwtClaims)
	if err != nil {
		return &auth.JWTClaims{}, fmt.Errorf("fail to generate token: %w", err)
	}

	return jwtClaims, nil
}
