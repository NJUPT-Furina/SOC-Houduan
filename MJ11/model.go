package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Count struct {
	CountName string
	password  string
}

func NewCount(countName, password string) (*Count, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("初始化密码失败: %v", err)
	}
	return &Count{
		CountName: countName,
		password:  string(hashedPassword),
	}, nil
}
