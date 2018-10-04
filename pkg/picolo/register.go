package picolo

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Register(clusterId string, nodeId string, nodeAddr string) error {
	client, err := fbApp.Firestore(ctx)
	if err != nil {
		fmt.Println("Error initializing database client:", err)
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
			fmt.Println("type conversion failed")
		}
		nodes = existingNodes
	}
	nodes[nodeId] = nodeAddr
	_, err = client.Collection("clusters").Doc(clusterId).Set(ctx, map[string]interface{}{
		"createdAt": firestore.ServerTimestamp,
		"nodes":     nodes,
	}, firestore.MergeAll)

	if err != nil {
		fmt.Println("Error registering node:", err)
		return err
	}
	return err
}
