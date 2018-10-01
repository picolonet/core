package main

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"os"
)

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile(os.Getenv("SERVICE_CREDS_FILE"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panicln("Error initializing app: %v\n", err)
		return nil
	}
	fbApp = app
	return app
}

func register(clusterId string) error {
	client, err := fbApp.Firestore(ctx)
	if err != nil {
		log.Panicln("Error initializing database client:", err)
		return err
	}
	defer client.Close()

	nodes := make(map[string]interface{})
	//check if cluster already exists and get nodes if it does
	dsnap, err := client.Collection("clusters").Doc(clusterId).Get(ctx)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			// do nothing, the cluster doesn't exist
		} else {
			fmt.Println("Error occured while fetching document:", err)
			return err
		}
	}
	if dsnap.Exists() {
		m := dsnap.Data()
		existingNodes, ok := m["nodes"].(map[string]interface{})
		if !ok {
			// Can't assert, handle error.
			log.Panicln("type conversion failed")
		}
		nodes = existingNodes
	}
	nodes[nodeId] = nodeAddr
	_, err = client.Collection("clusters").Doc(clusterId).Set(ctx, map[string]interface{}{
		"createdAt": firestore.ServerTimestamp,
		"nodes":     nodes,
	}, firestore.MergeAll)

	if err != nil {
		log.Panicln("Error registering node:", err)
		return err
	}
	return err
}
