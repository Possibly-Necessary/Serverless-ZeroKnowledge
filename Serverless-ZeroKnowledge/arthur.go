package main

import (
	"encoding/json"

	"github.com/0xdecaf/zkrp/bulletproofs"
	"github.com/nuclio/nuclio-sdk-go"
)

// Arthur the serverless function
func Arthur(context *nuclio.Context, event nuclio.Event) (interface{}, error) {

	// Decode the proof from JSON
	var decodedProof bulletproofs.ProofBPRP
	if err := json.Unmarshal(event.GetBody(), &decodedProof); err != nil {
		return nil, err
	}

	// Verify the proof
	ok, _ := decodedProof.Verify()

	if ok {
		context.Logger.Info("Arthur: Age verified to be in the range [18, 1000)")
	} else {
		context.Logger.Warn("Arthur: Age verification failed.")
	}

	return ok, nil
}
