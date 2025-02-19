package main

import (
	"fmt"
	"proto/auth"
)

func main() {
	req := &auth.AuthRequest{ // Используем структуру из protobuf
		Username: "test_user",
		Password: "secret",
	}

	fmt.Println("Auth request:", req)
}