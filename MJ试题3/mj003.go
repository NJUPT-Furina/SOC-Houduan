package main

import "fmt"

func main() {
	countName := "B26040721"
	name := "NJUPT-furina"
	password := "26040721"
	isNewuser := false //小目标1
	weekLogincount := []int{5, 7, 4, 1, 0, 7, 4}
	fmt.Printf("第1天的登录次数:%d\n", weekLogincount[0])
	fmt.Printf("第7天的登录次数:%d\n", weekLogincount[6])
	fmt.Printf("总共登陆了%d天\n", len(weekLogincount)) //小目标2
	weekLogincount[2] = 9
	fmt.Printf("第3天的登录次数:%d\n", weekLogincount[2]) //小目标3
	fmt.Printf("帐号卡信息：")
	fmt.Printf("用户名：%s\n", countName)
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("密码：%s\n", password)
	fmt.Printf("是否新用户：%t\n", isNewuser) //小目标4
}