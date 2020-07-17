package main

import (
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	validBody := `
		{
			"id": 1,
			"clubId": 1,
			"rankingCode": 1,
			"firstName": "Tacio",
			"lastName": "Costa",
			"address": "Somewhere",
			"phone": "12341234",
			"email": "tacio@email.com"
		}
	`
	invalidBody := `{"id": 0}`
	testCreateHandler(validBody, invalidBody, CreatePlayerHandler, t)
}
