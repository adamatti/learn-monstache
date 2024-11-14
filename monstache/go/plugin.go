package main

import (
	// "errors"
	"log"

	// "github.com/rwynn/monstache/monstachemap"
	"github.com/rwynn/monstache/v6/monstachemap"
	// "time"
)

func mapUser(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	log.Println("Processing user data")
	doc := input.Document
	/*
		session := input.Session
		db := session.DB(input.Database)
		locationCollection := db.C("location")
		locationId := doc["locationId"]

		if locationId != nil {
			doc["location"] = locationCollection.FindId(locationId)
		}
	*/

	delete(doc, "password")
	delete(doc, "version")

	output = &monstachemap.MapperPluginOutput{
		Document: doc,
		Index:    "users",
	}
	return
}

func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	if input.Collection == "users" {
		return mapUser(input)
	}

	doc := input.Document
	output = &monstachemap.MapperPluginOutput{Document: doc}
	return
}
