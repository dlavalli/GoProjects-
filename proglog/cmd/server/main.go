package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/dlavalli/GoProjects-/proglog/cmd/server/config"
	"github.com/dlavalli/GoProjects-/proglog/internal/server"
)

//go:embed resources/configs/proglog.json
var configResource []byte

func main() {
	var serverConfig = config.LoadServerConfiguration(configResource)
	fmt.Println(serverConfig.GetServerUrl())
	srv := server.NewHTTPServer(serverConfig.GetServerUrl())
	log.Fatal(srv.ListenAndServe())
}
