package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddress         string
	APIDomain             string
	APIPrefix             string
	EnableTLS             bool
	CertFile              string
	CertPassword          string
	DBDriver              string
	DBServer              string
	DBUser                string
	DBPassword            string
	DBPort                int
	DBName                string
	EncryptDataPassword   string
	EncryptConfigPassword string
}

func LoadConfig(sPath string) (*Config, error) {
	viper.AddConfigPath(sPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	// Read config
	config.ServerAddress = viper.GetString("server.address")
	config.APIDomain = viper.GetString("server.api_domain")
	config.APIPrefix = viper.GetString("server.api_prefix")
	config.EnableTLS = viper.GetBool("server.enable_tls")
	config.CertFile = viper.GetString("server.cert_file")
	config.CertPassword = viper.GetString("server.cert_password")
	config.DBDriver = viper.GetString("db.driver")
	config.DBServer = viper.GetString("db.server")
	config.DBUser = viper.GetString("db.user")
	config.DBPassword = viper.GetString("db.password")
	config.DBPort = viper.GetInt("db.port")
	config.DBName = viper.GetString("db.database")
	config.EncryptDataPassword = "US@PG#fu$CP%HM!2020&TMS"
	config.EncryptConfigPassword = "4asm$mcrt@198tqk$"
	return &config, nil
}
