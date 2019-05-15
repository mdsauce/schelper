package helper

import (
	"fmt"

	"github.com/mdsauce/schelper/logger"
)

func singleOutput(prob KnownProblem, logline []byte) {
	fmt.Printf("=======PROBLEM=======\n")
	logger.Disklog.Infof(`Problem Category: %s

Where: %s

Suggested Next Steps: 
%s

General Steps for this type of Disruption: 
%s

`, prob.Disruption.Category, logline, prob.NextSteps, prob.Disruption.GeneralSteps)
}

func metaOutput(metadata map[string]int) {
	logger.Disklog.Info("Metadata of All Problems Encountered")
	logger.Disklog.Info("------------------------------------")
	for key, val := range metadata {
		logger.Disklog.Infof("%s: %d", key, val)
	}
}

func lifecycleOutput(cycle [6]scLifecycle) {
	logger.Disklog.Info("Sauce Connect Lifecycle")
	logger.Disklog.Info("------------------------------------")
	for i := range cycle {
		if cycle[i].reached {
			logger.Disklog.Infof("%s reached", cycle[i].stage)
			logger.Disklog.Infof("-------------------------------------------")
		} else {
			logger.Disklog.Infof("%s *never* reached", cycle[i].stage)
			logger.Disklog.Infof("<----------------->")
		}
	}
}
