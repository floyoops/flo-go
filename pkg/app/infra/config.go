package infra

import "os"

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
	ContactEmailApp  string
}

func NewConfig(rootPath string) *Config {
	return &Config{
		RootPath:         rootPath,
		DatabaseHost:     getEnv("DATABASE_HOST", "localhost"),
		DatabasePort:     getEnv("DATABASE_PORT", "3006"),
		DatabaseName:     getEnv("DATABASE_NAME", "flogo"),
		DatabaseUser:     getEnv("DATABASE_USER", "root"),
		DatabasePassword: getEnv("DATABASE_PASSWORD", "toor"),
		SmtpHost:         getEnv("SMTP_HOST", "localhost"),
		SmtpPort:         getEnv("SMTP_PORT", "1025"),
		SmtpUsername:     getEnv("SMTP_USER", "flogo"),
		SmtpPassword:     getEnv("SMTP_PASSWORD", "toor"),
		ContactEmailApp:  getEnv("CONTACT_EMAIL_APP", "flogo@flogo.com"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
