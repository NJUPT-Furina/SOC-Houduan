package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Count struct {
	CountName string
	name      string
	password  string
	result    bool
}

func NewCount(countName, name, password string) (*Count, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("初始化密码失败: %v", err)
	}
	return &Count{
		CountName: countName,
		name:      name,
		password:  string(hashedPassword),
		result:    true,
	}, nil
}

func (c *Count) String() string {
	return fmt.Sprintf("账号名：%s/n昵称：%s/n密码摘要：%s/n创建结果：%t", c.CountName, c.name, c.password, c.result)
}
