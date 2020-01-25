package problems

import (
	"fmt"

	"github.com/mdsauce/schelper/logger"
)

type scLifecycle struct {
	stages [6]stage
}

type stage struct {
	name    string
	reached bool
	target  string
	line    int
	order   int
}

func (tun *scLifecycle) initLifecycle() {
	i := 0
	tun.stages[i].name = "Client attempted to start"
	tun.stages[i].target = "connecting to Sauce Labs REST API"
	tun.stages[i].order = i
	i++

	tun.stages[i].name = "Client started"
	tun.stages[i].target = "Started scproxy on port"
	tun.stages[i].order = i
	i++

	tun.stages[i].name = "Client attempted to connect to Maki"
	tun.stages[i].target = "connecting to tunnel VM"
	tun.stages[i].order = i
	i++

	tun.stages[i].name = "Sauce Connect Tunnel started"
	tun.stages[i].target = "Sauce Connect is up, you may start your tests."
	tun.stages[i].order = i
	i++

	tun.stages[i].name = "Client stopping attempts to reach Maki."
	tun.stages[i].target = "Connection closed"
	tun.stages[i].order = i
	i++

	tun.stages[i].name = "Client Shutdown"
	tun.stages[i].target = "Goodbye."
	tun.stages[i].order = i
	i++
}

func (tun scLifecycle) lifecycleOutput() {
	logger.Disklog.Info("Sauce Connect Lifecycle")
	logger.Disklog.Info("------------------------------------")
	for _, cycle := range tun.stages {
		if cycle.reached {
			logger.Disklog.Infof("Lifecycle Stage: %s", cycle.name)
			logger.Disklog.Infof("Found logline: '%s'", cycle.target)
			logger.Disklog.Infof("Found on logline %d", cycle.line)
		} else {
			logger.Disklog.Infof("------> Lifecycle Stage: %s not reached <------", cycle.name)
			logger.Disklog.Infof("Did not find logline: '%s'", cycle.target)
		}
		fmt.Println()
	}
}
