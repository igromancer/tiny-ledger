package main

import "github.com/igromancer/tiny-ledger/internal/server"

func main() {
	srv := server.Create()
	srv.Run()
}
