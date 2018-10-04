package picolo

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/jasonlvhit/gocron"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"time"
)

const version = "1.0.4"
const repo = "picolonet/core"
const selfUpdateTime = "13:00"
const selfUpdateTimeZone = "America/Los_Angeles"

func update() error {
	fmt.Println("Running self update")
	selfupdate.EnableLog()
	current := semver.MustParse(version)
	fmt.Println("Current version is", current)
	latest, err := selfupdate.UpdateSelf(current, repo)
	if err != nil {
		fmt.Println("Error self updating app:", err)
		return err
	}

	if current.Equals(latest.Version) {
		fmt.Println("Current binary is the latest version", version)
	} else {
		fmt.Println("Update successfully done to version", latest.Version)
		fmt.Println("Release notes:", latest.ReleaseNotes)
	}
	return nil
}

func ScheduleSelfUpdater() {
	PST, err := time.LoadLocation(selfUpdateTimeZone)
	if err != nil {
		fmt.Println(err)
		return
	}
	gocron.ChangeLoc(PST)
	gocron.Every(1).Day().At(selfUpdateTime).Do(update)
	<-gocron.Start()
}
