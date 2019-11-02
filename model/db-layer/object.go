package dblayer

import (
	"github.com/Myriad-Dreamin/minimum-template/config"
	"github.com/Myriad-Dreamin/minimum-template/types"
	"github.com/jinzhu/gorm"
	"time"
)

func wrapToObject(object interface{}, err error) (*Object, error) {
	return object.(*Object), err
}

func ObjectFactory() interface{} {
	return new(Object)
}

var (
	objectTraits = NewObjectTraits(Object{})
)

type Object struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
}

// TableName specification
func (Object) TableName() string {
	return "object"
}

func (Object) migrate() error {
	return objectTraits.Migrate()
}

func (d Object) GetID() uint {
	return d.ID
}

func (d *Object) Create() (int64, error) {
	return objectTraits.Create(d)
}

func (d *Object) Update() (int64, error) {
	return objectTraits.Update(d)
}

func (d *Object) UpdateFields(fields []string) (int64, error) {
	return objectTraits.UpdateFields(d, fields)
}

func (d *Object) Delete() (int64, error) {
	return objectTraits.Delete(d)
}

type ObjectDB struct{}

func NewObjectDB(logger types.Logger, _ *config.ServerConfig) (*ObjectDB, error) {
	return new(ObjectDB), nil
}

func GetObjectDB(logger types.Logger, _ *config.ServerConfig) (*ObjectDB, error) {
	return new(ObjectDB), nil
}

func (objectDB *ObjectDB) ID(id uint) (object *Object, err error) {
	return wrapToObject(objectTraits.ID(id))
}

type ObjectQuery struct {
	db *gorm.DB
}

func (objectDB *ObjectDB) QueryChain() *ObjectQuery {
	return &ObjectQuery{db: db}
}

func (objectDB *ObjectQuery) Order(order string, reorder ...bool) *ObjectQuery {
	objectDB.db = objectDB.db.Order(order, reorder...)
	return objectDB
}

func (objectDB *ObjectQuery) Page(page, pageSize int) *ObjectQuery {
	objectDB.db = objectDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return objectDB
}
func (objectDB *ObjectQuery) BeforeID(id uint) *ObjectQuery {
	objectDB.db = objectDB.db.Where("id <= ?", id)
	return objectDB
}

func (objectDB *ObjectQuery) Preload() *ObjectQuery {
	objectDB.db = objectDB.db.Set("gorm:auto_preload", true)
	return objectDB
}

func (objectDB *ObjectQuery) Query() (objects []Object, err error) {
	err = objectDB.db.Find(&objects).Error
	return
}
