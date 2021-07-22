package server

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/bagus-aulia/dot-test/app/pkg/models"
	"github.com/bagus-aulia/dot-test/app/server/book"
	"github.com/bagus-aulia/dot-test/app/server/borrow"
	"github.com/bagus-aulia/dot-test/app/server/middleware"
	"github.com/bagus-aulia/dot-test/app/server/user"
	"github.com/bagus-aulia/dot-test/config"
	"github.com/go-redis/redis"
	echo "github.com/labstack/echo/v4"
	colorable "github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB variable global db
	DB          *gorm.DB
	initialized = false
	redisDb     *redis.Client
)

func configureLogging() {
	lLevel := config.Get("server.log.level")
	fmt.Println("Setting log level to ", lLevel)
	switch strings.ToUpper(lLevel) {
	default:
		fmt.Println("Unknown level [", lLevel, "]. Log level set to ERROR")
		log.SetLevel(log.ErrorLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	}

	currentTime := time.Now()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/test_dot-" + currentTime.Format("2006-01-02") + ".log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     7, //days
		Level:      log.GetLevel(),
		Formatter: &log.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		log.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	log.SetLevel(log.GetLevel())
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	log.AddHook(rotateFileHook)
}

func recoverPanic() {
	if r := recover(); r != nil {
		log.WithField("panic", r).WithField("stack trace", string(debug.Stack())).Error("we panicked!")
	}
}

func dbMigration(DB *gorm.DB) {
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Borrow{})
	DB.AutoMigrate(&models.BorrowDetail{})
	DB.AutoMigrate(&models.User{})
}

// Start this server
func Start() {
	defer recoverPanic()
	configureLogging()

	// DB Connection
	dbHost := config.Get("db.host")
	dbPort := config.Get("db.port")
	dbUser := config.Get("db.user")
	dbPass := config.Get("db.password")
	dbName := config.Get("db.name")

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	var err error

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Auto migration
	dbMigration(DB)

	// Redis Connection
	redisDbConf := config.GetInt("redis.store")

	if s, err := strconv.Atoi(config.Get("redis.db")); err == nil {
		redisDbConf = s
	}

	redisDb = redis.NewClient(&redis.Options{
		Addr:     config.Get("redis.host") + ":" + config.Get("redis.port"),
		Password: "",          // no password set
		DB:       redisDbConf, // use default DB
	})

	e := echo.New()
	mdl := middleware.InitMiddleware()
	e.Use(mdl.CORS)
	timeoutContext := time.Duration(config.GetInt("server.timeout.read")) * time.Second

	// Server list
	book.Server(e, DB, redisDb, timeoutContext)
	borrow.Server(e, DB, redisDb, timeoutContext)
	user.Server(e, DB, redisDb, timeoutContext)

	e.Logger.Fatal(e.Start(config.Get("server.host") + config.Get("server.port")))

	fmt.Println("Server Run at ", config.Get("server.host"), config.Get("server.port"))
}
