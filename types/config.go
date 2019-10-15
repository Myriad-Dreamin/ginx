package types

type ConfigLoader func(cfg interface{}, cfgPath string) bool
