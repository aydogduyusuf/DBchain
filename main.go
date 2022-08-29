package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aydogduyusuf/DBchain/api"
	"github.com/aydogduyusuf/DBchain/blockchain"
	db "github.com/aydogduyusuf/DBchain/db/sqlc"
	"github.com/aydogduyusuf/DBchain/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect do db", err)
	}

	store := db.NewStore(conn)

	avalancheFuji := "https://api.avax-test.network/ext/bc/C/rpc"
	client := blockchain.InitNetwork(avalancheFuji)
	fmt.Printf("client: %v\n", client)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server")
	}

	
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}