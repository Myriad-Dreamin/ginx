package dblayer

import (
	"github.com/Myriad-Dreamin/minimum-lib/crypto"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	userTraits            Traits
	userQueryNameFunc     where1Func
	userQueryNickNameFunc where1Func
	userQueryPhoneFunc    where1Func
)

func injectUserTraits() error {
	userTraits = NewTraits(User{})

	userQueryNameFunc = userTraits.Where1("name = ?")
	userQueryNickNameFunc = userTraits.Where1("nick_name = ?")
	userQueryPhoneFunc = userTraits.Where1("phone = ?")
	return nil
}

func wrapToUser(user interface{}, err error) (*User, error) {
	if user == nil {
		return nil, err
	}
	return user.(*User), err
}

type User struct {
	ID        uint      `dorm:"id" gorm:"column:id;primary_key;not_null"`
	CreatedAt time.Time `dorm:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time `dorm:"updated_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"`
	LastLogin time.Time `dorm:"last_login" gorm:"column:last_login;default:CURRENT_TIMESTAMP;not null;" json:"last_login"`

	NickName string `dorm:"nick_name" gorm:"column:nick_name;unique;not_null"`
	Name     string `dorm:"name" gorm:"column:name;not_null"`
	Password string `dorm:"password" gorm:"column:password;not_null"`
	Phone    string `dorm:"phone" gorm:"column:phone;unique;not_null"`
}

// TableName specification
func (User) TableName() string {
	return "user"
}

func (User) migrate() error {
	return userTraits.Migrate()
}

func (d User) GetID() uint {
	return d.ID
}

func (d *User) Create() (int64, error) {
	return userTraits.Create(d)
}

func (d *User) Update() (int64, error) {
	return userTraits.Update(d)
}

func (d *User) UpdateFields(fields []string) (int64, error) {
	return userTraits.UpdateFields(d, fields)
}

func (d *User) Delete() (int64, error) {
	return userTraits.Delete(d)
}

func (d *User) Register() (int64, error) {
	var err error
	d.Password, err = crypto.NewPasswordString(d.Password)
	if err != nil {
		return 0, err
	}

	return d.Create()
}

func (d *User) ResetPassword(password string) (int64, error) {
	var err error
	d.Password, err = crypto.NewPasswordString(password)
	if err != nil {
		return 0, err
	}

	return d.UpdateFields([]string{"password"})
}

func (d *User) AuthenticatePassword(pswd string) (bool, error) {
	return crypto.CheckPasswordString(pswd, d.Password)
}

type UserDB struct{}

func NewUserDB(_ module.Module) (*UserDB, error) {
	return new(UserDB), nil
}

func GetUserDB(_ module.Module) (*UserDB, error) {
	return new(UserDB), nil
}

func (userDB *UserDB) Filter(f *Filter) (user []User, err error) {
	err = userTraits.Filter(f, &user)
	return
}

func (userDB *UserDB) FilterI(f interface{}) (interface{}, error) {
	return userDB.Filter(f.(*Filter))
}

func (userDB *UserDB) ID(id uint) (user *User, err error) {
	return wrapToUser(userTraits.ID(id))
}

type UserQuery struct {
	db *gorm.DB
}

func (userDB *UserDB) QueryChain() *UserQuery {
	return &UserQuery{db: p.GormDB}
}

func (userDB *UserQuery) Order(order string, reorder ...bool) *UserQuery {
	userDB.db = userDB.db.Order(order, reorder...)
	return userDB
}

func (userDB *UserQuery) Page(page, pageSize int) *UserQuery {
	userDB.db = userDB.db.Limit(pageSize).Offset((page - 1) * pageSize)
	return userDB
}
func (userDB *UserQuery) BeforeID(id uint) *UserQuery {
	userDB.db = userDB.db.Where("id <= ?", id)
	return userDB
}

func (userDB *UserQuery) Preload() *UserQuery {
	userDB.db = userDB.db.Set("gorm:auto_preload", true)
	return userDB
}

func (userDB *UserQuery) Query() (users []User, err error) {
	err = userDB.db.Find(&users).Error
	return
}

func (userDB *UserDB) Has(id uint) (has bool, err error) {
	return userTraits.Has(id)
}

func (userDB *UserDB) Query(id uint) (user *User, err error) {
	return wrapToUser(userTraits.ID(id))
}

func (userDB *UserDB) QueryName(id string) (user *User, err error) {
	return wrapToUser(userQueryNameFunc(id))
}

func (userDB *UserDB) QueryNickName(id string) (user *User, err error) {
	return wrapToUser(userQueryNickNameFunc(id))
}

func (userDB *UserDB) QueryPhone(id string) (user *User, err error) {
	return wrapToUser(userQueryPhoneFunc(id))
}
