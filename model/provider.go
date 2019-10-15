package model

import splayer "github.com/Myriad-Dreamin/ginx/model/sp-layer"

type Provider = splayer.Provider

func NewProvider(namespace string) *Provider {
	return splayer.NewProvider(namespace)
}

func SetProvider(p *Provider) *Provider {
	return splayer.SetProvider(p)
}
