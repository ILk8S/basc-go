package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)


const (
	emailRegexExpPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)
type UserHander struct {
	emailRegexExp *regexp.Regexp
	passwordRegexExp *regexp.Regexp
}

func NewUserHandler() *UserHander {
	return &UserHander{
		emailRegexExp: regexp.MustCompile(emailRegexExpPattern, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (h *UserHander) RegisterRoutes(server *gin.Engine) {

}

func (h *UserHander) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email string `json:"email"`
		Password string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
}

func (h *UserHander) Login(ctx *gin.Context) {

}

func (h *UserHander) Edit(ctx *gin.Context) {

}

func (h *UserHander) Profile(ctx *gin.Context) {

}