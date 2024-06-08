package config

import (
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/labstack/gommon/log"
	"os"
)

type Config struct {
	RootPath         string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	SmtpHost         string
	SmtpPort         string
	SmtpUsername     string
	SmtpPassword     string
	ContactEmailApp  *model.Email
}

func NewConfig(rootPath string) *Config {
	contactEmailApp, err := model.NewEmail(getEnv("CONTACT_EMAIL_APP", "flogo@flogo.com"))
	if err != nil {
		log.Fatalf("error for key env CONTACT_EMAIL_APP: %v", err)
	}

	return &Config{
		RootPath:         rootPath,
		DatabaseHost:     getEnv("DATABASE_HOST", "localhost"),
		DatabasePort:     getEnv("DATABASE_PORT", "3306"),
		DatabaseName:     getEnv("DATABASE_NAME", "flogo"),
		DatabaseUser:     getEnv("DATABASE_USER", "root"),
		DatabasePassword: getEnv("DATABASE_PASSWORD", "toor"),
		SmtpHost:         getEnv("SMTP_HOST", "localhost"),
		SmtpPort:         getEnv("SMTP_PORT", "1025"),
		SmtpUsername:     getEnv("SMTP_USER", "flogo"),
		SmtpPassword:     getEnv("SMTP_PASSWORD", "toor"),
		ContactEmailApp:  contactEmailApp,
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
