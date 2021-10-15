package auth

import (
	"banking_crowd/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Role struct {
	Roles string
}

func Permission(r *Role) gin.HandlerFunc {
	return func(c *gin.Context) {

		role := c.MustGet("role").(string)

		success := checkRole(r.Roles, role)
		if !success {
			response := helper.APIResponse("You dont have permission", http.StatusMethodNotAllowed, "error", nil)
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, response)
			return
		}

	}
}

func checkRole(role string, currentRole string) bool {
	return strings.Contains(role, currentRole)
}
