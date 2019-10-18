package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/jinzhu/gorm"
	"time"
)

var submissionModel *dorm.Model

type Submission struct {
	ID uint `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (Submission) TableName() string {
	return "submission"
}

func (Submission) migrate() error {
	err := db.AutoMigrate(&Submission{}).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	submissionModel, err = dormDB.Model(&Submission{})
	return err
}

func (d Submission) GetID() uint {
	return d.ID
}

func (d *Submission) Create() (int64, error) {
	rdb := db.Create(d)
	return rdb.RowsAffected, rdb.Error
}

func (d *Submission) Update() (int64, error) {
	rdb := db.Save(d)
	return rdb.RowsAffected, rdb.Error
}

func (d *Submission) UpdateFields(fields []string) (int64, error) {
	return submissionModel.Anchor(d).Select(fields...).UpdateFields()
}

func (d *Submission) Delete() (int64, error) {
	rdb := db.Delete(d)
	return rdb.RowsAffected, rdb.Error
}


type SubmissionDB struct {}

func NewSubmissionDB(logger types.Logger) (*SubmissionDB, error) {
	return new(SubmissionDB), nil
}

func GetSubmissionDB(logger types.Logger) (*SubmissionDB, error) {
	return new(SubmissionDB), nil
}


func (objDB *SubmissionDB) ID(id uint) (obj *Submission, err error) {
	obj = new(Submission)
	rdb := db.First(obj, id)
	err = rdb.Error
	if rdb.RecordNotFound() {
		obj = nil
		err = nil
	}
	return
}

type SubmissionQuery struct {
	db *gorm.DB
}

func (objDB *SubmissionDB) QueryChain() *SubmissionQuery {
	return &SubmissionQuery{db:db}
}

func (objDB *SubmissionQuery) Order(order string, reorder ...bool) *SubmissionQuery {
	objDB.db = objDB.db.Order(order, reorder...)
	return objDB
}

func (objDB *SubmissionQuery) Page(page, pageSize int) *SubmissionQuery {
	objDB.db = objDB.db.Limit(pageSize).Offset((page-1)*pageSize)
	return objDB
}
func (objDB *SubmissionQuery) BeforeID(id uint) *SubmissionQuery {
	objDB.db = objDB.db.Where("id <= ?", id)
	return objDB
}

func (objDB *SubmissionQuery) Preload() *SubmissionQuery {
	objDB.db = objDB.db.Set("gorm:auto_preload", true)
	return objDB
}

func (objDB *SubmissionQuery) Query() (objs []Submission, err error) {
	err  = objDB.db.Find(&objs).Error
	return
}

