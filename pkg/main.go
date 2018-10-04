package main

import (
	"fmt"
	"github.com/picolonet/core/pkg/picolo"
	log "github.com/sirupsen/logrus"
)

//todo: fix logging

// todo: change these
const nodeId = "test"
const nodeAddr = "testIpAddr"
const clusterId = "clusterId"

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	go picolo.ScheduleSelfUpdater()
	picolo.InitAppWithServiceAccount()
	picolo.Register(clusterId, nodeId, nodeAddr)
	picolo.ThrowFlare(nodeId)

	fmt.Println("Started node")
}
