// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package trialtbl allow to use a Unit Test Suite Table Framework.

This is made with the concept of Trial, Experiment and Suite. The Suite
is a set of tests called Experiment, that can run one of more
verifications and analysis called Trial. That being said, the Suite is
made receiving a table of Experiments and its Trials. After creation
the Suite can be put to test.

Test of Suite it's made receiving a testing.T variable, and a test 
function that represents the Test to Run on each Experiment. The Suite
calls for each Experiment, the function test.

Test function given to Suite.Test must receive an Experiment as
parameter, which should be used to register Result for Trials. The Test
function can register as many Results as the developer finds necessary,
but the Results will only get calculated if assigns to a valid Trial.

When the test function is done for the Experiment, Suite run for each
Trial on that experiment, a comparison between values expected and
registered values on that Trial. If the values doesn't match, the Test
will fail, printing the error as the Result registered told to print.

Usage

The package can be used like this, importing testing:

	type SumOper interface {
		Sum(int, int) int
		LastResultAsString() string
	}

	// type SumOp struct {
	//		...
	// }

	// Define choosen trials, expected result and factors used.
	trialtbl.NewSuite(
		trialtbl.NewExperiment(	
			trialtbl.NewTrial(true, 1, 2, 3),
			trialtbl.NewTrial("3"),
		),
		trialtbl.NewExperiment(	
			trialtbl.NewTrial(true, 2, 3, 5),
			trialtbl.NewTrial("5"),
		),
		trialtbl.NewExperiment(	
			trialtbl.NewTrial(true, -3, 2, -1),
			trialtbl.NewTrial("-1"),
		),
		trialtbl.NewExperiment(	
			trialtbl.NewTrial(true, 2, 0, 2),
			trialtbl.NewTrial("2"),
		),
	).Test(t, func(e *trialtbl.Experiment){
		// Execute Experiments using the same function.
		var s SumOp

		// Verify Sum return value.
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			val := s.Sum(f[0].(int), f[1].(int))
			sig := "s.Sum(\"%s\", \"%s\")"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Verify Sum string format.
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			val := s.LastResultAsString()
			sig := "s.LastResultAsString()"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
*/
package trialtbl