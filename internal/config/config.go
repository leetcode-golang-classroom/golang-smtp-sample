package config

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-smtp-sample/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	Port          int    `mapstructure:"PORT"`
	FromEmail     string `mapstructure:"FROM_EMAIL"`
	FromEmailSmtp string `mapstructure:"FROM_EMAIL_SMTP"`
	SmtpAddr      string `mapstructure:"SMTP_ADDR"`
	SmtpSecret    string `mapstructure:"SMTP_SECRET"`
}

var AppConfig *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.SetConfigName(".env")
	v.AutomaticEnv()
	util.FailOnError(v.BindEnv("PORT"), "failed to bind env PORT")
	util.FailOnError(v.BindEnv("FROM_EMAIL"), "failed to bind env FROM_EMAIL")
	util.FailOnError(v.BindEnv("FROM_EMAIL_SMTP"), "failed to bind env FROM_EMAIL_SMTP")
	util.FailOnError(v.BindEnv("SMTP_ADDR"), "failed to bind env SMTP_ADDR")
	util.FailOnError(v.BindEnv("SMTP_SECRET"), "failed to bind env SMTP_SECRET")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("load from environmentt variable")
	}
	err = v.Unmarshal(&AppConfig)
	if err != nil {
		util.FailOnError(err, "failed to read environment")
	}
}
