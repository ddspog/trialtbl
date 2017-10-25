package trialtbl

// Trial contains variables for verifications on tests. The idea is
// that each verification, or Trial will contain Factors to call on a
// test function, a value that is expected to be returned as Result, and
// some other attributes setted by Experiment methods.
type Trial struct {
	Factors []interface{}
	Expected interface{}
	result interface{}
	signature string
}

// NewTrial creates a new Trial with the needed information, the factors
// to call on the test method, and the value expected to receive as 
// result.
func NewTrial(r interface{}, f ...interface{}) (t *Trial) {
	t = &Trial{
		Factors: f,
		Expected: r,
	}
	return
}