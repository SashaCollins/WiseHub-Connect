package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"github.com/joho/godotenv"
)

// type global
type singleton map[string]string

type gitHubConfig struct {
	Username string
	APIToken string
}
type droneCIConfig struct {
	Host string
	APIToken string
}
type herokuConfig struct {
	Username string
	Password string
	APIToken string
}
type configuration struct {
	GitHub gitHubConfig
	DroneCI droneCIConfig
	Heroku herokuConfig
	DebugMode bool
	//UserRoles []string
	//MaxUsers  int
}

var (
	lock sync.Mutex
	once   sync.Once
	config *configuration
)
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
// New returns a new Config struct
func GetConfig() *configuration {
	once.Do(func() {
		config = &configuration{
			GitHub: gitHubConfig{
				Username: getEnv("GITHUB_USERNAME", ""),
				APIToken: getEnv("GITHUB_API_TOKEN", ""),
			},
			DroneCI: droneCIConfig{
				Host: getEnv("DRONE_HOST", ""),
				APIToken: getEnv("DRONE_API_TOKEN", ""),
			},
			Heroku: herokuConfig{
				Username: getEnv("HEROKU_USERNAME", ""),
				Password: getEnv("HEROKU_PASSWORD", ""),
				APIToken: getEnv("HEROKU_API_TOKEN", ""),
			},
			DebugMode: getEnvAsBool("DEBUG_MODE", true),
			//UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),
			//MaxUsers:  getEnvAsInt("MAX_USERS", 1),
		}
	})
	fmt.Println(config)
	return config
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	//lock.Lock()
	//defer lock.Unlock()
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
//func getEnvAsInt(name string, defaultVal int) int {
//	valueStr := getEnv(name, "")
//	if value, err := strconv.Atoi(valueStr); err == nil {
//		return value
//	}
//
//	return defaultVal
//}

// Helper to read an environment variable into a string slice or return default value
//func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
//	valStr := getEnv(name, "")
//
//	if valStr == "" {
//		return defaultVal
//	}
//
//	val := strings.Split(valStr, sep)
//
//	return val
//}