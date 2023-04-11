package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DahuiSheng/gobackend/ent"
	_ "github.com/lib/pq"
)

// 変数の定義：名前と型
var (
	client *ent.Client
	err    error
)

func init() {
	client, err = ent.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	fmt.Println("接続しました")
	defer client.Close()
}

// 外部からアクセスしたいデータなので、先頭文字を大文字にする
// DBを読み出すため。
func GetClient() *ent.Client {
	return client
}
