package router

type NopValidator struct{}

func (NopValidator) Enforce(...interface{}) (bool, error) {
	return true, nil
}
