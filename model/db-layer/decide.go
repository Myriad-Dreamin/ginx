package dblayer

import (
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
)

type Traits = traits.ModelTraits
type Interface = traits.Interface
type TraitsAcceptObject = traits.ORMObject

func NewTraits(t TraitsAcceptObject) Traits {
	tt := traits.NewModelTraits(t, db, dormDB)
	return tt
}


var (
	_                = "TraitsDefinitionHook"
	NewObjectTraits  = NewTraits
	NewUserTraits    = NewTraits
)
