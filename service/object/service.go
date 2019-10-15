package objectservice

import (
	"github.com/Myriad-Dreamin/ginx/model"
	"github.com/Myriad-Dreamin/ginx/types"
)

type Service struct {
	db *model.ObjectDB
	logger types.Logger
}

func NewService(logger types.Logger, provider *model.Provider) (a *Service, err error) {
	a = new(Service)
	a.db = provider.ObjectDB()
	a.logger = logger
	return
}

/*
type Object struct {
}
*/


