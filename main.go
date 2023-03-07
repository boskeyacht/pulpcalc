package main

import (
	"fmt"

	cmd "github.com/baribari2/pulp-calculator/cmd/pulpcalc"
	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/sashabaranov/go-openai"
	"github.com/sausheong/goreplicate"
)

// Confidence
// 1-50% Possible
// 51-98% Almost Certain
// 99% Certain
func main() {
	cfg := types.InitConfig()

	// if cfg.DictEndpoint == "" {
	// 	cfg.DictEndpoint = "dict://dict.dict.org"
	// } else {
	// 	s, err := dict.Dial("tcp", cfg.DictEndpoint)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	cfg.DictEndpoint= s
	// }

	driver, err := neo4j.NewDriver(cfg.NeoEndpoint, neo4j.BasicAuth(cfg.NeoUser, cfg.NeoPassword, ""))
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

	model := goreplicate.NewModel("openai", "whisper", "30414ee7c4fffc37e260fcab7842b5be470b9b840f2b608f5baa9bbef9a259ed")
	client := goreplicate.NewClient(cfg.ReplicateKey, model)
	cfg.ReplicateClient = client

	oc := openai.NewClient(cfg.OpenAIKey)
	cfg.OpenAIClient = oc

	cmd.Execute(cfg)
}
