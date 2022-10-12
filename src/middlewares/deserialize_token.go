package middlewares

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeserializeUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("Authorization")
		if len(accessToken) == 0 {
			c.Next()
			return
		}

		refreshToken := c.Request.Header.Get("X-Refresh")
		if len(refreshToken) == 0 {
			c.Next()
			return
		}

		accessClaim, err := utils.DecodeToken(accessToken)

		if err != nil && len(refreshToken) > 0 {
			newAccessToken, user := utils.ReIssueAccessToken(db, refreshToken)

			if user != nil {
				c.Header("X-Access", newAccessToken)
				c.Set("User", user)
			}

			c.Next()
			return
		}

		data := accessClaim["data"].(map[string]interface{})

		user := &entities.User{
			Id:        data["id"].(string),
			Username:  data["username"].(string),
			Email:     data["email"].(string),
			Name:      data["name"].(string),
			CreatedAt: helpers.ConvertFloat64ToInt64(data["created_at"].(float64)),
			UpdatedAt: helpers.ConvertFloat64ToInt64(data["updated_at"].(float64)),
		}

		c.Set("User", user)
		c.Next()
	}
}
