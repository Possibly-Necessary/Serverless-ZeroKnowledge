package main

import (
    "context"
    "encoding/json"
    "math/big"
    "github.com/0xdecaf/zkrp/bulletproofs"
    "github.com/nuclio/nuclio-sdk-go"
)

// Merlin the serverless function
func Merlin(context *nuclio.Context, event nuclio.Event)(interface{}, error){
	
	// Set up the range, [18, 1000) in this case
	params, _ := bulletproofs.SetupGeneric(18, 1000)

	// Merlin's age is 666
	context.Logger.Info("Merlin's age is 666 (his secret).")
	bigSecret := new(big.Int).SetInt64(int64(666))

	// Create the zero-knowledge range proof
	proof, _ := bulletproofs.ProveGeneric(bigSecret, params)

	// Encode the proof to JSON
	jsonEncoded, _ := json.Marshal(proof)

	return nuclio.Response{
		StatusCode: 200,
		ContentType: "application/json",
		Body: jsonEncoded,
	}, nil
}