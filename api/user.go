package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kadycui/gin-rank/database"
	"github.com/kadycui/gin-rank/middleware"
	"github.com/kadycui/gin-rank/model"
	"github.com/kadycui/gin-rank/serialize"
	"github.com/kadycui/gin-rank/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @BasePath		/api/v1
// @Summary		用户注册
// @Description	新用户注册
// @Tags			用户
// @Version		v1.0
// @Accept			json
// @Produce		json
// @Param			name		query		string	false	"用户名"
// @Param			telephone	query		string	true	"电话 "
// @Param			password	query		string	true	"密码 "
// @Success		200			{string}	string	"注册成功!"
// @Router			/api/v1/auth/register  [post]
func Register(ctx *gin.Context) {

	DB := database.GetDB()
	// 获取参数
	name := ctx.Request.FormValue("name")
	telephone := ctx.Request.FormValue("telephone")
	password := ctx.Request.FormValue("password")

	log.Println("请求参数:", name, telephone, password)

	//数据验证
	if len(telephone) != 11 {
		serialize.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}
	if len(password) < 6 {
		serialize.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}

	// 如果没有传名字,返回一个随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		serialize.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已存在")
		return
	}

	// 创建用户
	hashPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		serialize.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return

	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassWord),
	}
	DB.Create(&newUser)

	// 返回结果
	serialize.Response(ctx, http.StatusOK, 200, nil, "注册成功")

}

type User struct {
	Name     string `json:"name"`
	Password int64  `json:"password"`
}

// @BasePath		/api/v1
// @Summary		用户登录
// @Description	新用户注册获取鉴权信息
// @Tags			用户
// @Version		v1.0
// @Accept			json
// @Produce		json
// @Param			name		query		string	false	"用户名"
// @Param			password	query		string	true	"密码 "
// @Success		200			{string}	string	"登录成功!"
// @Router			/api/v1/auth/login  [post]
func Login(ctx *gin.Context) {
	DB := database.GetDB()
	// 获取参数
	username := ctx.Request.FormValue("name")
	password := ctx.Request.FormValue("password")

	fmt.Println(username, password)
	utils.LogEntry.Info(username, password)

	// 判断用户是否存在
	var user model.User
	DB.Where("name = ?", username).First(&user)
	if user.ID == 0 {
		serialize.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return

	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		serialize.Response(ctx, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}

	// 发放Token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		serialize.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "系统异常")

		log.Printf("token generate error: %v\n", err)
		return
	}

	// 返回结果
	// 返回结果
	serialize.Response(ctx, http.StatusOK, 0, gin.H{"token": token}, "登录成功!")
	//ctx.JSON(200, gin.H{
	//	"code": 200,
	//	"msg":  "登录成功!",
	//	"data": gin.H{"token": token},
	//})

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	serialize.Response(ctx, http.StatusOK, 200, gin.H{
		"user": serialize.ToUserSerialize(user.(model.User)),
	}, "")

	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{
	//		"user": serialize.ToUserSerialize(user.(model.User)),
	//	},
	//})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
