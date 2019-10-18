package splayer

import (
	"github.com/Myriad-Dreamin/ginx/config"
	dblayer "github.com/Myriad-Dreamin/ginx/model/db-layer"
	"github.com/Myriad-Dreamin/ginx/types"
)

type Submission = dblayer.Submission

//struct {
//	dblayer.Submission
//}

//func (c *Submission) Update() (int64, error) {
//	return c.Submission.Update()
//}

//type _userConverter struct {}
//var SubmissionConverter _userConverter
//func (_userConverter) convert(obj dblayer.Submission) Submission { return Submission{Submission: obj} }
//func (_userConverter) reconvert(obj Submission) dblayer.Submission { return obj.Submission }
//func (_userConverter) converts(objs []dblayer.Submission) []Submission { return *(*[]Submission)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconverts(objs []Submission) []dblayer.Submission { return *(*[]dblayer.Submission)(unsafe.Pointer(&objs)) }
//func (_userConverter) convertP(obj *dblayer.Submission) *Submission { return (*Submission)(unsafe.Pointer(obj)) }
//func (_userConverter) reconvertP(obj *Submission) *dblayer.Submission { return (*dblayer.Submission)(unsafe.Pointer(obj)) }
//func (_userConverter) convertPs(objs []*dblayer.Submission) []*Submission { return *(*[]*Submission)(unsafe.Pointer(&objs)) }
//func (_userConverter) reconvertPs(objs []*Submission) []*dblayer.Submission { return *(*[]*dblayer.Submission)(unsafe.Pointer(&objs)) }

type SubmissionDB struct {
	dblayer.SubmissionDB
}

func NewSubmissionDB(logger types.Logger, _ *config.ServerConfig) (*SubmissionDB, error) {
	cdb, err := dblayer.NewSubmissionDB(logger)
	if err != nil {
		return nil, err
	}
	db := new(SubmissionDB)
	db.SubmissionDB = *cdb
	return db, nil
}

func GetSubmissionDB(logger types.Logger, _ *config.ServerConfig) (*SubmissionDB, error) {
	cdb, err := dblayer.GetSubmissionDB(logger)
	if err != nil {
		return nil, err
	}
	db := new(SubmissionDB)
	db.SubmissionDB = *cdb
	return db, nil
}

func (s *Provider) SubmissionDB() *SubmissionDB {
	return s.submissionDB
}
