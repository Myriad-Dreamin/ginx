package control

import (
	"github.com/Myriad-Dreamin/go-model-traits/gorm-crud-dao"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-template/model/db-layer"
	"time"
)

type UserService interface {
	UserServiceSignatureXXX() interface{}
	ListUsers(c controller.MContext)
	Login(c controller.MContext)
	Register(c controller.MContext)
	ChangePassword(c controller.MContext)
	InspectUser(c controller.MContext)
	GetUser(c controller.MContext)
	PutUser(c controller.MContext)
	Delete(c controller.MContext)
}

type ListUsersRequest = gorm_crud_dao.Filter

type ListUsersReply struct {
	Code  int             `json:"code"form:"code"`
	Users []ListUserReply `json:"users"form:"users"`
}

func PSerializeListUsersReply(_code int, _users []ListUserReply) *ListUsersReply {
	return &ListUsersReply{
		Code:  _code,
		Users: _users,
	}
}

func SerializeListUsersReply(_code int, _users []ListUserReply) ListUsersReply {
	return ListUsersReply{
		Code:  _code,
		Users: _users,
	}
}

func PackSerializeListUsersReply(_code []int, _users [][]ListUserReply) (pack []ListUsersReply) {
	for i := range _code {
		pack = append(pack, SerializeListUsersReply(_code[i], _users[i]))
	}
	return
}

type ListUserReply struct {
	NickName  string    `json:"nick_name"form:"nick_name"`
	LastLogin time.Time `json:"last_login"form:"last_login"`
}

func PSerializeListUserReply(vUser dblayer.User) *ListUserReply {
	return &ListUserReply{
		NickName:  vUser.NickName,
		LastLogin: vUser.LastLogin,
	}
}

func SerializeListUserReply(vUser dblayer.User) ListUserReply {
	return ListUserReply{
		NickName:  vUser.NickName,
		LastLogin: vUser.LastLogin,
	}
}

func PackSerializeListUserReply(vUser []dblayer.User) (pack []ListUserReply) {
	for i := range vUser {
		pack = append(pack, SerializeListUserReply(vUser[i]))
	}
	return
}

type LoginRequest struct {
	Id       uint   `json:"id"form:"id"`
	NickName string `json:"nick_name"form:"nick_name"`
	Phone    string `form:"phone"json:"phone"`
	Password string `json:"password"form:"password"binding:"required"`
}

func PSerializeLoginRequest(user *dblayer.User, _password string) *LoginRequest {
	return &LoginRequest{
		Id:       user.ID,
		NickName: user.NickName,
		Phone:    user.Phone,
		Password: _password,
	}
}

func SerializeLoginRequest(user *dblayer.User, _password string) LoginRequest {
	return LoginRequest{
		Id:       user.ID,
		NickName: user.NickName,
		Phone:    user.Phone,
		Password: _password,
	}
}

func PackSerializeLoginRequest(user []*dblayer.User, _password []string) (pack []LoginRequest) {
	for i := range user {
		pack = append(pack, SerializeLoginRequest(user[i], _password[i]))
	}
	return
}

type LoginReply struct {
	Code         int      `form:"code"json:"code"`
	Id           uint     `json:"id"form:"id"`
	Identity     []string `json:"identity"form:"identity"`
	Phone        string   `form:"phone"json:"phone"`
	NickName     string   `form:"nick_name"json:"nick_name"`
	Name         string   `json:"name"form:"name"`
	Token        string   `json:"token"form:"token"`
	RefreshToken string   `json:"refresh_token"form:"refresh_token"`
}

func PSerializeLoginReply(_code int, user *dblayer.User, _identity []string, _token string, _refreshToken string) *LoginReply {
	return &LoginReply{
		Code:         _code,
		Id:           user.ID,
		Identity:     _identity,
		Phone:        user.Phone,
		NickName:     user.NickName,
		Name:         user.Name,
		Token:        _token,
		RefreshToken: _refreshToken,
	}
}

func SerializeLoginReply(_code int, user *dblayer.User, _identity []string, _token string, _refreshToken string) LoginReply {
	return LoginReply{
		Code:         _code,
		Id:           user.ID,
		Identity:     _identity,
		Phone:        user.Phone,
		NickName:     user.NickName,
		Name:         user.Name,
		Token:        _token,
		RefreshToken: _refreshToken,
	}
}

func PackSerializeLoginReply(_code []int, user []*dblayer.User, _identity [][]string, _token []string, _refreshToken []string) (pack []LoginReply) {
	for i := range _code {
		pack = append(pack, SerializeLoginReply(_code[i], user[i], _identity[i], _token[i], _refreshToken[i]))
	}
	return
}

type RegisterRequest struct {
	Name     string `json:"name"form:"name"binding:"required"`
	Password string `json:"password"form:"password"binding:"required"`
	NickName string `json:"nick_name"form:"nick_name"binding:"required"`
	Phone    string `binding:"required"json:"phone"form:"phone"`
}

func PSerializeRegisterRequest(_name string, _password string, _nickName string, _phone string) *RegisterRequest {
	return &RegisterRequest{
		Name:     _name,
		Password: _password,
		NickName: _nickName,
		Phone:    _phone,
	}
}

func SerializeRegisterRequest(_name string, _password string, _nickName string, _phone string) RegisterRequest {
	return RegisterRequest{
		Name:     _name,
		Password: _password,
		NickName: _nickName,
		Phone:    _phone,
	}
}

func PackSerializeRegisterRequest(_name []string, _password []string, _nickName []string, _phone []string) (pack []RegisterRequest) {
	for i := range _name {
		pack = append(pack, SerializeRegisterRequest(_name[i], _password[i], _nickName[i], _phone[i]))
	}
	return
}

type RegisterReply struct {
	Code int  `json:"code"form:"code"`
	Id   uint `json:"id"form:"id"`
}

func (r RegisterReply) GetID() uint {
	return r.Id
}

func PSerializeRegisterReply(_code int, user *dblayer.User) *RegisterReply {
	return &RegisterReply{
		Code: _code,
		Id:   user.ID,
	}
}

func SerializeRegisterReply(_code int, user *dblayer.User) RegisterReply {
	return RegisterReply{
		Code: _code,
		Id:   user.ID,
	}
}

func PackSerializeRegisterReply(_code []int, user []*dblayer.User) (pack []RegisterReply) {
	for i := range _code {
		pack = append(pack, SerializeRegisterReply(_code[i], user[i]))
	}
	return
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"form:"old_password"binding:"required"`
	NewPassword string `json:"new_password"form:"new_password"binding:"required"`
}

func PSerializeChangePasswordRequest(_oldPassword string, _newPassword string) *ChangePasswordRequest {
	return &ChangePasswordRequest{
		OldPassword: _oldPassword,
		NewPassword: _newPassword,
	}
}

func SerializeChangePasswordRequest(_oldPassword string, _newPassword string) ChangePasswordRequest {
	return ChangePasswordRequest{
		OldPassword: _oldPassword,
		NewPassword: _newPassword,
	}
}

func PackSerializeChangePasswordRequest(_oldPassword []string, _newPassword []string) (pack []ChangePasswordRequest) {
	for i := range _oldPassword {
		pack = append(pack, SerializeChangePasswordRequest(_oldPassword[i], _newPassword[i]))
	}
	return
}

type InspectUserReply struct {
	Code int           `json:"code"form:"code"`
	User *dblayer.User `json:"user"form:"user"`
}

func PSerializeInspectUserReply(_code int, _user *dblayer.User) *InspectUserReply {
	return &InspectUserReply{
		Code: _code,
		User: _user,
	}
}

func SerializeInspectUserReply(_code int, _user *dblayer.User) InspectUserReply {
	return InspectUserReply{
		Code: _code,
		User: _user,
	}
}

func PackSerializeInspectUserReply(_code []int, _user []*dblayer.User) (pack []InspectUserReply) {
	for i := range _code {
		pack = append(pack, SerializeInspectUserReply(_code[i], _user[i]))
	}
	return
}

type GetUserReply struct {
	Code      int       `json:"code"form:"code"`
	NickName  string    `json:"nick_name"form:"nick_name"`
	LastLogin time.Time `json:"last_login"form:"last_login"`
}

func PSerializeGetUserReply(_code int, user *dblayer.User) *GetUserReply {
	return &GetUserReply{
		Code:      _code,
		NickName:  user.NickName,
		LastLogin: user.LastLogin,
	}
}

func SerializeGetUserReply(_code int, user *dblayer.User) GetUserReply {
	return GetUserReply{
		Code:      _code,
		NickName:  user.NickName,
		LastLogin: user.LastLogin,
	}
}

func PackSerializeGetUserReply(_code []int, user []*dblayer.User) (pack []GetUserReply) {
	for i := range _code {
		pack = append(pack, SerializeGetUserReply(_code[i], user[i]))
	}
	return
}

type PutUserRequest struct {
	Phone string `json:"phone"form:"phone"`
}

func PSerializePutUserRequest(user *dblayer.User) *PutUserRequest {
	return &PutUserRequest{
		Phone: user.Phone,
	}
}

func SerializePutUserRequest(user *dblayer.User) PutUserRequest {
	return PutUserRequest{
		Phone: user.Phone,
	}
}

func PackSerializePutUserRequest(user []*dblayer.User) (pack []PutUserRequest) {
	for i := range user {
		pack = append(pack, SerializePutUserRequest(user[i]))
	}
	return
}
