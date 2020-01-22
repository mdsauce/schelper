package problems

import (
	"fmt"

	"github.com/mdsauce/schelper/logger"
)

func singleOutput(prob KnownProblem, logline []byte) {
	fmt.Printf("=======PROBLEM=======\n")
	logger.Disklog.Infof(`Problem: %s

Where: %s

Suggested Next Steps: 
%s

`, prob.Name, logline, prob.NextSteps)
}

func noMakiReplyOutput(reply bool) {
	if !reply {
		fmt.Printf("=======PROBLEM=======\n")
		logger.Disklog.Infof(`Problem: No Reply from the Allocated Maki

Suggested Next Steps:
No reply was heard from the Maki.  i.e. never found 000000000001.  This is usually the result of a bad connection or a lack of access to *.miso.saucelabs.com.  Attempt to cURL a maki like 'curl maki_SOME_NUMBER.miso.saucelabs.com. Obtain the maki number from the latest sauce connect logs.

`)
	}
}

func problemsOutput(problemsdata map[string]int) {
	logger.Disklog.Info("Metadata of All Problems Encountered")
	logger.Disklog.Info("------------------------------------")
	for key, val := range problemsdata {
		logger.Disklog.Infof("%s: %d", key, val)
	}
	fmt.Println()
}

func lifecycleOutput(cycle [6]scLifecycle) {
	logger.Disklog.Info("Sauce Connect Lifecycle")
	logger.Disklog.Info("------------------------------------")
	for i := range cycle {
		if cycle[i].reached {
			logger.Disklog.Infof("Lifecycle Stage: %s", cycle[i].stage)
			logger.Disklog.Infof("Found logline: '%s'", cycle[i].target)
			logger.Disklog.Infof("Found on logline %d", cycle[i].line)
		} else {
			logger.Disklog.Infof("------> Lifecycle Stage: %s not reached <------", cycle[i].stage)
			logger.Disklog.Infof("Did not find logline: '%s'", cycle[i].target)
		}
		fmt.Println()
	}
}
