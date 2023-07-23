package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kadycui/gin-rank/database"
	"github.com/kadycui/gin-rank/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取 authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足!",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足!",
			})
			ctx.Abort()
			return
		}

		// 获取claim
		userId := claims.UserId
		DB := database.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足!",
			})
			ctx.Abort()
			return
		}

		// 用户存在, 讲用户信息写入上下文
		ctx.Set("user", user)

		ctx.Next()

	}

}
