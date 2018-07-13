package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

func main() {
	jobs := sdk.Jobs{
		sdk.Job{
			Handler: func() error {
				log.Println("Pipeline log output.")
				return nil
			},
			Title:       "First pipeline task",
			Description: "Just some pipeline test task.",
			Priority:    0,
		},
	}

	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
