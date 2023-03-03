package main

import (
	"fmt"

	cmd "github.com/baribari2/pulp-calculator/cmd/pulpcalc"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Confidence
// 1-50% Possible
// 51-98% Almost Certain
// 99% Certain
func main() {
	cfg := initConfig()

	// if DictServer == "" {
	// 	DictServer = "dict://dict.dict.org"
	// } else {
	// 	s, err := dict.Dial("tcp", DictServer)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	cfg.DictServer = s
	// }

	driver, err := neo4j.NewDriver(NeoEndpoint, neo4j.BasicAuth(NeoUser, NeoPassword, ""))
	if err != nil {
		fmt.Println(err)

		return
	}

	err = driver.VerifyConnectivity()
	if err != nil {
		fmt.Println(err)

		return
	}

	s := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	cfg.Neo4j = s

	cmd.Execute(cfg)
}
