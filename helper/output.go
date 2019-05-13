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
