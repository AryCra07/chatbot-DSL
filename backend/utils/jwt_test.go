package utils

import (
	"testing"
)

func TestJWT(t *testing.T) {
	// 测试生成token
	userId := "testUser"
	password := "testPassword"
	token, err := GenerateToken(userId, password)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}

	// 测试解析token
	parsedToken, claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("Error parsing token: %v", err)
	}

	// 检查解析后的claims是否正确
	if claims.UserId != userId {
		t.Errorf("Expected UserId to be %s, got %s", userId, claims.UserId)
	}

	if claims.Password != password {
		t.Errorf("Expected Password to be %s, got %s", password, claims.Password)
	}

	// 检查token是否有效
	if !parsedToken.Valid {
		t.Error("Parsed token is not valid")
	}
}
