package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/blang/semver"
	"github.com/jasonlvhit/gocron"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"google.golang.org/genproto/googleapis/type/latlng"
	"log"
)

//todo: fix logging

const version = "1.0.4"

// todo: change these
const nodeId = "test"
const nodeAddr = "testIpAddr"

var location = &latlng.LatLng{Latitude: 9, Longitude: 179}

var fbApp *firebase.App
var ctx = context.Background()

func main() {
	gocron.Every(1).Day().At("13:00").Do(selfUpdate) //todo: check timezone
	initializeAppWithServiceAccount()
	register("clusterId")
	throwFlare()
	<-gocron.Start()
}

func selfUpdate() error {
	selfupdate.EnableLog()

	current := semver.MustParse(version)
	fmt.Println("Current version is", current)
	latest, err := selfupdate.UpdateSelf(current, "picolonet/core")
	if err != nil {
		log.Panicln("Error self updating app: %v\n", err)
		return err
	}

	if current.Equals(latest.Version) {
		fmt.Println("Current binary is the latest version", version)
	} else {
		fmt.Println("Update successfully done to version", latest.Version)
		fmt.Println("Release notes:\n", latest.ReleaseNotes)
	}
	return nil
}
