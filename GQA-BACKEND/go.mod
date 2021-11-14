module gin-quasar-admin

go 1.17

require (
	github.com/casbin/gorm-adapter/v3 v3.4.4
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-gonic/gin v1.7.4
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/mojocn/base64Captcha v1.3.5
	github.com/spf13/viper v1.9.0
	go.uber.org/zap v1.19.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.22.3
)

require (
	github.com/casbin/casbin/v2 v2.37.4
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa // indirect
	gorm.io/driver/postgres v1.2.2 // indirect
	gorm.io/driver/sqlserver v1.2.1 // indirect
)
