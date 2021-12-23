//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/masseelch/elk"
)

func main() {
	ex, err := elk.NewExtension(
		elk.GenerateSpec("openapi.json"),
		elk.GenerateHandlers(),
	)
	if err != nil {
		log.Fatalf("creating elk extension: %+v", err)
	}

	opts := []entc.Option{
		entc.FeatureNames("privacy"),
		entc.FeatureNames("sql/upsert"),
		entc.FeatureNames("schema/snapshot"),
		entc.Extensions(ex),
	}
	err = entc.Generate("./schema", &gen.Config{}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %+v", err)
	}
}
