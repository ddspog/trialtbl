package trialtbl

import (
	"fmt"
	"testing"
)

// TestSumOper is a test interface to suite test a struct.
type TestSumOper interface {
	Sum(int, int) int
	LastResultAsString() string
}

// TestSumOper is a test structure to suite test a struct.
type TestSumOp struct {
	lastResultAsString string
}

// Sum is a test method to suite test a struct.
func (s *TestSumOp) Sum(a, b int) (x int) {
	x = a + b
	s.lastResultAsString = fmt.Sprintf("%v", x)
	return
}

// LastResultAsString is a test method to suite test a struct.
func (s *TestSumOp) LastResultAsString() (r string) {
	r = s.lastResultAsString
	return
}

// Test if a suite can run a set of experiments.
func TestSuiteTest(t *testing.T) {
	// Define choosen trials, expected result and factors used.
	NewSuite(
		NewExperiment(	
			NewTrial(true, 1, 2, 3),
			NewTrial("3"),
		),
		NewExperiment(	
			NewTrial(true, 2, 3, 5),
			NewTrial("5"),
		),
		NewExperiment(	
			NewTrial(true, -3, 2, -1),
			NewTrial("-1"),
		),
		NewExperiment(	
			NewTrial(true, 2, 0, 2),
			NewTrial("2"),
		),
	).Test(t, func(e *Experiment){
		// Execute Experiments using the same function.
		var s TestSumOp

		// Verify Sum return value.
		e.RegisterResult(0, func(f ...interface{}) (r *Result) {
			val := s.Sum(f[0].(int), f[1].(int)) == f[2].(int)
			sig := "s.Sum(\"%v\", \"%v\") == \"%v\""
			r = NewResult(val, sig)
			return
		})

		// Verify Sum string format.
		e.RegisterResult(1, func(f ...interface{}) (r *Result) {
			val := s.LastResultAsString()
			sig := "s.LastResultAsString()"
			r = NewResult(val, sig)
			return
		})
	})
}