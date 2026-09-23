package main

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Count struct {
	CountName string
	password  string
	status    string
}

func NewCount(countName, password string) (*Count, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("初始化密码失败: %v", err)
	}
	return &Count{
		CountName: countName,
		password:  string(hashedPassword),
		status:    "active",
	}, nil
}

func (c *Count) UpdatePassword(newPassword string) error {
	if c == nil {
		return errors.New("无法更新密码: 账户未初始化")
	}
	if c.password == "" {
		return errors.New("无法更新密码: 当前密码为空")
	}
	if len(newPassword) < 6 {
		return errors.New("新密码长度必须至少为6个字符")
	}
	if newPassword == c.password {
		return errors.New("新密码不能与当前密码相同")
	}
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("更新密码失败: %w", err)
	}
	c.password = string(newHashedPassword)
	return nil
}

func validateCount(status string) {
	if status == "blocked" {
		fmt.Println("账号风险拦截")
		os.Exit(1)
	}
}
