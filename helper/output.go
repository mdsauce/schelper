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
	fmt.Printf("\nMetadata of All Problems Encountered\n")
	fmt.Printf("------------------------------------\n")
	for key, val := range metadata {
		logger.Disklog.Infof("%s: %d", key, val)
	}
}
