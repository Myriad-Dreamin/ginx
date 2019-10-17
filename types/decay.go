package types

type DecayResult struct {
	First interface{}
	Err   error
}

func Decay(d interface{}, e error) DecayResult {
	return DecayResult{d, e}
}

