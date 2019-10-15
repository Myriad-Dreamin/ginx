package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/ginx/types"
	"github.com/jinzhu/gorm"
	"time"
)

var objectModel *dorm.Model

type Object struct {
	ID uint `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (Object) TableName() string {
	return "obj"
}

func (Object) migrate() error {
	err := db.AutoMigrate(&Object{}).Error
	if err != nil {
		return err
	}

	//db.AddIndex()
	objectModel, err = dormDB.Model(&Object{})
	return err
}

func (d Object) GetID() uint {
	return d.ID
}

func (d *Object) Create() (int64, error) {
	rdb := db.Create(d)
	return rdb.RowsAffected, rdb.Error
}

func (d *Object) Update() (int64, error) {
	rdb := db.Save(d)
	return rdb.RowsAffected, rdb.Error
}

func (d *Object) UpdateFields(fields []string) (int64, error) {
	return objectModel.Anchor(d).Select(fields...).UpdateFields()
}

func (d *Object) Delete() (int64, error) {
	rdb := db.Delete(d)
	return rdb.RowsAffected, rdb.Error
}


type ObjectDB struct {}

func NewObjectDB(logger types.Logger) (*ObjectDB, error) {
	return new(ObjectDB), nil
}

func GetObjectDB(logger types.Logger) (*ObjectDB, error) {
	return new(ObjectDB), nil
}


func (objDB *ObjectDB) ID(id uint) (obj *Object, err error) {
	obj = new(Object)
	rdb := db.First(obj, id)
	err = rdb.Error
	if rdb.RecordNotFound() {
		obj = nil
		err = nil
	}
	return
}

type ObjectQuery struct {
	db *gorm.DB
}

func (objDB *ObjectDB) QueryChain() *ObjectQuery {
	return &ObjectQuery{db:db}
}

func (objDB *ObjectQuery) Order(order string, reorder ...bool) *ObjectQuery {
	objDB.db = objDB.db.Order(order, reorder...)
	return objDB
}

func (objDB *ObjectQuery) Page(page, pageSize int) *ObjectQuery {
	objDB.db = objDB.db.Limit(pageSize).Offset((page-1)*pageSize)
	return objDB
}
func (objDB *ObjectQuery) BeforeID(id uint) *ObjectQuery {
	objDB.db = objDB.db.Where("id <= ?", id)
	return objDB
}

func (objDB *ObjectQuery) Preload() *ObjectQuery {
	objDB.db = objDB.db.Set("gorm:auto_preload", true)
	return objDB
}

func (objDB *ObjectQuery) Query() (objs []Object, err error) {
	err  = objDB.db.Find(&objs).Error
	return
}

