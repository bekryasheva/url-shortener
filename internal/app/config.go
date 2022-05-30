package app

import (
	"log"

	"github.com/spf13/viper"
)

type SectionAPI struct {
	Address string `mapstructure:"address"`
}

type SectionURL struct {
	AddressPrefix string `mapstructure:"address_prefix"`
}

type SectionPostgresql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Sslmode  string `mapstructure:"sslmode"`
}

type Config struct {
	API      SectionAPI        `mapstructure:"api"`
	URL      SectionURL        `mapstructure:"url"`
	Storage  string            `mapstructure:"storage"`
	Database SectionPostgresql `mapstructure:"postgres"`
}

func ReadConfigFromFile(path string) (Config, error) {
	config := &Config{
		API: SectionAPI{
			Address: ":8080",
		},
		URL: SectionURL{
			AddressPrefix: "http://localhost:8080/",
		},
		Storage: "postgresql",
		Database: SectionPostgresql{
			Host:     "localhost",
			Port:     "5432",
			User:     "urlshortener",
			Password: "password",
			Name:     "urlshortener",
			Sslmode:  "disable",
		},
	}

	viper.SetConfigFile(path)
	viper.BindEnv("postgres.host", "DB_HOST")
	viper.BindEnv("postgres.port", "DB_PORT")
	viper.BindEnv("postgres.user", "DB_USER")
	viper.BindEnv("postgres.password", "DB_PASSWORD")
	viper.BindEnv("postgres.name", "DB_NAME")
	viper.BindEnv("postgres.sslmode", "DB_SSLMODE")

	viper.BindEnv("storage", "STORAGE")
	viper.BindEnv("api.address", "API_ADDRESS")

	viper.BindEnv("url.address", "URL_ADDRESS_PREFIX")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("%v\n", err)
		return Config{}, err
	}

	viper.GetString("api.port")

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return *config, nil
}
