// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

// Card defines model for Card.
type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// DeckReturnCard2Params defines parameters for DeckReturnCard2.
type DeckReturnCard2Params struct {

	// Short-form or long-form encoding of the card to return to the deck
	Card *string `json:"card,omitempty"`
}

// DeckReturnCardJSONBody defines parameters for DeckReturnCard.
type DeckReturnCardJSONBody Card

// DeckReturnCardJSONRequestBody defines body for DeckReturnCard for application/json ContentType.
type DeckReturnCardJSONRequestBody DeckReturnCardJSONBody
