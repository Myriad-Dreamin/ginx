package control

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

var AuthCates []interface{}

// @Category Auth - Group Grant/Revoke/Check Api Group
// @Description Grant Api Group
// @Path /v1/auth/group/user/:id
type authGrantApiGroupCate interface {
}

// @Category Auth - Sugar - Group Admin Grant/Revoke/Check Api Group
// @Description Grant Api Group
// @Path /v1/auth/sugar/group/admin/user/:id
type authGrantAdminApiGroupCate interface {
}

func init() {
	var (
		a authGrantApiGroupCate      = 0
		b authGrantAdminApiGroupCate = 0
	)
	AuthCates = []interface{}{
		&a,
		&b,
	}
}

/* auth
 * refresh token GET: 刷新登陆用token
 */
type AuthService interface {
	AuthSignatureXXX() interface{}
	// /refresh-token GET
	RefreshToken(c controller.MContext)
}
