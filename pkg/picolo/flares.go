package picolo

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"google.golang.org/genproto/googleapis/type/latlng"
)

var location = &latlng.LatLng{Latitude: 9, Longitude: 179}

func ThrowFlare(nodeId string) error {
	client, err := fbApp.Firestore(ctx)
	if err != nil {
		fmt.Println("Error initializing database client:", err)
		return err
	}
	defer client.Close()

	_, err = client.Collection("flares").Doc(nodeId).Set(ctx, map[string]interface{}{
		"lastFired": firestore.ServerTimestamp,
		"location":  location,
	}, firestore.MergeAll)

	if err != nil {
		fmt.Printf("Error throwing flare: %v\n", err)
		return err
	}
	return err
}
