package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/jasonlvhit/gocron"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "1.0.3"

func selfUpdate() error {
	selfupdate.EnableLog()

	current := semver.MustParse(version)
	fmt.Println("Current version is", current)
	latest, err := selfupdate.UpdateSelf(current, "picolonet/core")
	if err != nil {
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

func main() {
	gocron.Every(1).Day().At("13:00").Do(selfUpdate)
	<-gocron.Start()
}
