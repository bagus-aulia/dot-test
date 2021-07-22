package config

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defCfg      map[string]string
	initialized = false
)

// initialize this configuration
func initialize() {
	viper.SetConfigFile(".env")
	viper.SetEnvPrefix("dot")
	viper.ReadInConfig()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	defCfg = make(map[string]string)

	defCfg["server.log.level"] = viper.GetString(`DOT_SERVER_LOG_LEVEL`) // valid values are trace, debug, info, warn, error, fatal

	defCfg["db.type"] = "postgres"
	defCfg["server.address"] = fmt.Sprintf("%s:%s", viper.GetString(`DOT_SERVER_HOST`), viper.GetString(`DOT_SERVER_PORT`))
	defCfg["server.host"] = viper.GetString(`DOT_SERVER_HOST`)
	defCfg["server.port"] = viper.GetString(`DOT_SERVER_PORT`)
	defCfg["server.timeout.read"] = viper.GetString(`DOT_SERVER_TIMEOUT_READ`)

	defCfg["db.host"] = viper.GetString(`DOT_DB_HOST`)
	defCfg["db.port"] = viper.GetString(`DOT_DB_PORT`)
	defCfg["db.user"] = viper.GetString(`DOT_DB_USER`)
	defCfg["db.password"] = viper.GetString(`DOT_DB_PASSWORD`)
	defCfg["db.name"] = viper.GetString(`DOT_DB_NAME`)
	defCfg["db.maxidle"] = viper.GetString(`DOT_DB_MAXIDLE`)
	defCfg["db.maxopen"] = viper.GetString(`DOT_DB_MAXOPEN`)

	defCfg["redis.host"] = viper.GetString(`DOT_REDIS_HOST`)
	defCfg["redis.port"] = viper.GetString(`DOT_REDIS_PORT`)
	defCfg["redis.password"] = viper.GetString(`DOT_REDIS_PASSWORD`)
	defCfg["redis.db"] = viper.GetString(`DOT_REDIS_DB`)
	defCfg["redis.store"] = viper.GetString(`DOT_REDIS_STORE`)
	// defCfg["server.host"] = "localhost"
	// defCfg["server.port"] = ":3000"
	// defCfg["server.timeout.read"] = "300"

	// defCfg["db.host"] = "localhost"
	// defCfg["db.port"] = "5432"
	// defCfg["db.user"] = "postgres"
	// defCfg["db.password"] = "postgres"
	// defCfg["db.name"] = "test-dot"
	// defCfg["db.maxidle"] = "3"
	// defCfg["db.maxopen"] = "10"

	// defCfg["redis.host"] = "localhost"
	// defCfg["redis.port"] = "6379"
	// defCfg["redis.password"] = ""
	// defCfg["redis.db"] = "test-dot"
	// defCfg["redis.store"] = "0"

	for k := range defCfg {
		err := viper.BindEnv(k)
		if err != nil {
			log.Errorf("Failed to bind env \"%s\" into configuration. Got %s", k, err)
		}
	}

	initialized = true
}

// SetConfig put configuration key value
func SetConfig(key, value string) {
	viper.Set(key, value)
}

// Get fetch configuration as string value
func Get(key string) string {
	if !initialized {
		initialize()
	}
	ret := viper.GetString(key)
	if len(ret) == 0 {
		if ret, ok := defCfg[key]; ok {
			return ret
		}
		log.Debugf("%s config key not found", key)
	}
	return ret
}

// Set configuration key value
func Set(key, value string) {
	defCfg[key] = value
}

// GetInt fetch configuration as integer value
func GetInt(key string) int {
	if len(Get(key)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(Get(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}
