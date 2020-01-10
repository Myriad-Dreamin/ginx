package router

import (
	"github.com/Myriad-Dreamin/minimum-template/control/auth"
)

func ApplyAuth(router *RootRouter) {
	// var agi = router.AuthApiRouter.Group
	// agi.RevokeGroup.Use(agi.Auth.AdminOnly())
	// agi.GrantGroup.Use(agi.Auth.AdminOnly())

	// var aggMap = router.AuthApiRouter.Sugar.DynamicList
	// for _, agg := range aggMap {
	// 	agg.Revoke.Use(agg.Auth.AdminOnly())
	// 	agg.Grant.Use(agg.Auth.AdminOnly())
	// }
	//agi.CheckGroup.Use(agi.Auth.AdminOnly())

	var uig = router.UserRouter.IDRouter
	uig.ChangePassword.Use(uig.Auth.Build(auth.UserEntity.Write()))
	uig.Put.Use(uig.Auth.Build(auth.UserEntity.Write()))
	uig.Delete.Use(uig.Auth.AdminOnly())

}
