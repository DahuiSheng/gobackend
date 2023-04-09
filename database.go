package main

import (
	"log"

	"github.com/DahuiSheng/gobackend/ent"
)

// 変数の定義：名前と型
var (
	client *ent.Client
	err    error
)

func init() {
	client, err = ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
}

// 外部からアクセスしたいデータなので、先頭文字を大文字にする
// DBを読み出すため。
func GetClient() *ent.Client {
	return client
}
