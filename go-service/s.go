package main

import (
	"fmt"

	"github.com/Wertygool/WertyTGNotTradeBot/protos/auth"
)

func main() {
	req := &auth.AuthRequest{ // Используем структуру из protobuf
		Username: "test_user",
		Password: "secret",
	}

	fmt.Println("Auth request:", req)
}