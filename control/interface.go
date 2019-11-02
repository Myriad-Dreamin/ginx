package control

import "github.com/gin-gonic/gin"

//go:generate package-attach-to-path -generate_register_map


/* auth
 * refresh token GET: 刷新登陆用token
 */
type Auth interface {
	// /refresh-token GET
	RefreshToken(c *gin.Context)
}

