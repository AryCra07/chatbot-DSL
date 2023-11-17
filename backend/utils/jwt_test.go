package utils

import (
	"backend/consts"
	"backend/log"
	"testing"
)

const (
	userId   string = "userIdForTest"
	password string = "veryStrongPassword"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(userId, password)
	if err != nil {
		return
	}
	log.Info(consts.Test, token)
}

func TestParseToken(t *testing.T) {
	token, err := GenerateToken(userId, password)
	if err != nil {
		return
	}
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJkNDY0YzMxMi04NTFiLTExZWUtYjVmMi0wMGZmMDZlNzMxYzUiLCJwYXNzd29yZCI6ImZjMzA1ZmE5M2Q4YzExNTVhMmZhOGNlMGQwNGI3NTFhIiwiZXhwIjoxNzAwMjIxMDczLCJpYXQiOjE3MDAyMDY2NzMsImlzcyI6IkFyeUNyYTA3Iiwic3ViIjoidXNlciB0b2tlbiJ9.EuK-SMygq_NxeqY-NQXu0VDmN1MWYWCn53hAaN6IctY"
	tokenParse, claims, err := ParseToken(token)
	if tokenParse.Valid {
		log.Info(consts.Test, claims.UserId)
	} else {
		log.Info(consts.Test, "token is invalid:"+claims.UserId)
	}
}
