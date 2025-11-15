package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ILk8S/basc-go/internal/repository"
	"github.com/ILk8S/basc-go/internal/repository/dao"
	"github.com/ILk8S/basc-go/internal/service"
	"github.com/ILk8S/basc-go/internal/web"
	"github.com/ILk8S/basc-go/internal/web/middleware"
	"github.com/ILk8S/basc-go/pkg/ginx/middleware/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	server := initWebServer()
	initUserHdl(db, server)
	server.Run(":8899")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.1.40:3306)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		//AllowHeaders:     []string{"Content-Type"},
		AllowHeaders:  []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "/users")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		fmt.Println("Middleware")
	})
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.40:6379",
		Password: "pA3L48JS",
	})
	//限制每秒最多1个请求
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 1).Build())
	useJWT(server)
	return server
}

func useJWT(server *gin.Engine) {
	login := &middleware.LoginJWTMiddlewareBuilder{}
	server.Use(login.CheckLogin())
}

func useSession(server *gin.Engine) {
	login := &middleware.LoginMiddlewareBuilder{}
	store := cookie.NewStore([]byte("secret"))
	//store, err := redis.NewStore(16, "tcp",
	//	"192.168.1.40:6379", "",
	//	"pA3L48JS", []byte("JXMYyZfJI9kFrzq0ggoHmSqE6UAUuW8L"))
	//if err != nil {
	//	panic(errors.New("redis connection faild"))
	//}
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}
