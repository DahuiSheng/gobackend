package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込みできませんでした: %v", err)
	}

}
