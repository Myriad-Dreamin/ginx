package fcg

type MaybeInitializer func() error

func Calls(funcs []MaybeInitializer) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
