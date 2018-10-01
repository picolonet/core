package main

import (
	"cloud.google.com/go/firestore"
	"log"
)

func throwFlare() error {
	client, err := fbApp.Firestore(ctx)
	if err != nil {
		log.Panicln("Error initializing database client:", err)
		return err
	}
	defer client.Close()

	_, err = client.Collection("flares").Doc(nodeId).Set(ctx, map[string]interface{}{
		"lastFired": firestore.ServerTimestamp,
		"location":  location,
	}, firestore.MergeAll)

	if err != nil {
		log.Panicln("Error throwing flare: %v\n", err)
		return err
	}
	return err
}
