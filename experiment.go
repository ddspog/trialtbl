package trialtbl

// Experiment stores a set of Trials to be verified when the Experiment
// it's run on Suite.
type Experiment struct {
	Trials []*Trial
}

// NewExperiment creates a new Experiment, with the set of Trials to be
// verified when Experiment it's run.
func NewExperiment(t ...*Trial) (e *Experiment) {
	e = &Experiment{
		Trials: t,
	}
	return
}

// Result it's a auxiliary type to contain return values for the Trials.
type Result struct {
	Signature string
	Value interface{}
}

// NewResult creates a Result. This function it's supposed to be used on
// return of RegisterResult functions. Put the value found, and a string
// explaining what was called to found the value.
func NewResult(v interface{}, s string) (r *Result) {
	r = &Result{
		Signature: s,
		Value: v,
	}
	return
}

// RegisterResult calculate f Result for trial i, if it exists on 
// Experiment.
func (e *Experiment) RegisterResult(i int, f func(...interface{}) *Result) {
	if i < len(e.Trials) {
		r := f(e.Trials[i].Factors...)
		e.Trials[i].result = r.Value
		e.Trials[i].signature = r.Signature
	}
}