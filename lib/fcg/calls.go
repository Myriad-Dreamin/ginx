package fcg

type MaybeInitor func() error

func Calls(funcs []MaybeInitor) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
