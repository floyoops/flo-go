package config

import (
	"fmt"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	RootPath         string
	ServerPortHttp   int
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
	HttpAllowOrigins []string
}

func NewConfig() (*Config, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	contactEmailApp, err := model.NewEmail(getEnv("CONTACT_EMAIL_APP", "flogo@flogo.com"))
	if err != nil {
		return nil, fmt.Errorf("error for key env CONTACT_EMAIL_APP:: %w", err)
	}
	err = godotenv.Load(string(rootPath) + "/.env")
	if err != nil {
		return nil, fmt.Errorf("error on load .env:: %w", err)
	}

	return &Config{
		RootPath:         rootPath,
		ServerPortHttp:   getEnvAsInt("SERVER_HTTP_PORT", 8080),
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
		HttpAllowOrigins: getEnvAsSlice("HTTP_ALLOW_ORIGINS", []string{"http://localhost:3000"}, ","),
	}, nil
}

func (c *Config) GetDatabaseDns() string {
	return c.DatabaseUser + ":" + c.DatabasePassword + "@tcp(" + c.DatabaseHost + ":" + c.DatabasePort + ")/" + c.DatabaseName
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	val := strings.Split(valStr, sep)
	return val
}
