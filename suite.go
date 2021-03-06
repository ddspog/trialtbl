package trialtbl

import (
	"fmt"
	"testing"
)

// Suite it's the table of experiments to run on tests.
type Suite struct {
	Experiments []*Experiment 
}

// NewSuite create a new Suite from a set of Experiments with
// validation on Trials.
func NewSuite(e ...*Experiment) (s *Suite) {
	s = &Suite{
		Experiments: e,
	}
	return
}

// Test execute the function exp for each Experiment on Suite, and then
// compares Trials values with Results on each Experiment. When finding
// a mismatched value for a Trial, it reports an error following the
// signature and factors on the Trial failed.
func (s *Suite) Test(t *testing.T, exp func(*Experiment)) {
	for ie, e := range s.Experiments {
		exp(e)
		for it, trl := range e.Trials {
			if trl.result != trl.Expected {
				var sigfmt string
				if len(trl.Factors) > 0 {
					sigfmt = fmt.Sprintf(trl.signature, trl.Factors...)
				} else {
					sigfmt = fmt.Sprintf(trl.signature)
				}
				logh := fmt.Sprintf("\tExperiment %v, Trial %v\n", ie, it)
				logv := fmt.Sprintf("\t%s: %v\n", sigfmt, trl.result)
				loge := fmt.Sprintf("\tExpected: %v\n", trl.Expected)
				t.Error(logh + logv + loge)
			}
		}
	}
}