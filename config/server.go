package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	DBUsername          = "DB_USERNAME"
	DBPassword          = "DB_PASSWORD"
	DBHost              = "DB_HOST"
	DBPort              = "DB_PORT"
	DBName              = "DB_DATABASE"
	DBURL               = "DB_URL"
	EnvApp              = "ENV_APP"
	LogStdout           = "LOG_ENABLE_STDOUT"
	LogFile             = "LOG_ENABLE_FILE"
	KCClientId          = "KC_CLIENT_ID"
	KCClientSecret      = "KC_CLIENT_SECRET"
	KCRealm             = "KC_REALM"
	KCBasePath          = "KC_BASE_PATH"
	KCHttpManagmentPath = "KC_HTTP_MANAGEMENT_PATH"
	RedisAddr           = "REDIS_ADDRESS"
	S3AccessKey         = "S3_ACCESS_KEY"
	S3SecretKey         = "S3_SECRET_KEY"
	S3RegionName        = "S3_REGION_NAME"
	S3Endpoint          = "S3_ENDPOINT"
	S3Bucket            = "S3_BUCKET"
	TimeZone            = "TIMEZONE"
	SwaggerEnv          = "SWAGGER_ENV"
	PublicKeycloakURL   = "PUBLIC_KEYCLOAK_URL"
)

type Server struct {
	DBUsername          string
	DBPassword          string
	DBHost              string
	DBPort              int
	DBName              string
	DBURL               string
	EnvApp              string
	EnableStdout        bool
	EnableFile          bool
	KCClientId          string
	KCClientSecret      string
	KCRealm             string
	KCBasePath          string
	KCHttpManagmentPath string
	RedisAddr           string
	S3AccessKey         string
	S3SecretKey         string
	S3RegionName        string
	S3Endpoint          string
	S3Bucket            string
	TimeZone            string
	SwaggerEnv          string
	PublicKeycloakURL   string
}

func NewServer() *Server {
	viper.AutomaticEnv()

	viper.SetDefault(DBPort, 5432)
	viper.SetDefault(LogStdout, true)
	viper.SetDefault(LogFile, false)

	viper.SetDefault(TimeZone, "Asia/Ho_Chi_Minh")
	if viper.GetString(S3Bucket) == "" {
		log.Print("You need to set S3_BUCKET environment variable")
	}

	return &Server{
		DBUsername:          viper.GetString(DBUsername),
		DBPassword:          viper.GetString(DBPassword),
		DBHost:              viper.GetString(DBHost),
		DBPort:              viper.GetInt(DBPort),
		DBName:              viper.GetString(DBName),
		DBURL:               viper.GetString(DBURL),
		EnvApp:              viper.GetString(EnvApp),
		EnableStdout:        viper.GetBool(LogStdout),
		EnableFile:          viper.GetBool(LogFile),
		KCClientId:          viper.GetString(KCClientId),
		KCClientSecret:      viper.GetString(KCClientSecret),
		KCRealm:             viper.GetString(KCRealm),
		KCBasePath:          viper.GetString(KCBasePath),
		KCHttpManagmentPath: viper.GetString(KCHttpManagmentPath),
		RedisAddr:           viper.GetString(RedisAddr),
		S3AccessKey:         viper.GetString(S3AccessKey),
		S3SecretKey:         viper.GetString(S3SecretKey),
		S3RegionName:        viper.GetString(S3RegionName),
		S3Endpoint:          viper.GetString(S3Endpoint),
		S3Bucket:            viper.GetString(S3Bucket),
		TimeZone:            viper.GetString(TimeZone),
		SwaggerEnv:          viper.GetString(SwaggerEnv),
		PublicKeycloakURL:   viper.GetString(PublicKeycloakURL),
	}
}
