package config

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	General             GeneralConfig
	SQL                 SQLConfig
	ApplicationMessages map[string]ApplicationMessages
	CurrentLanguage     string
	AccessTokenSecret   string
	RefreshTokenSecret  string
	SendGridAPIKey      string
	SendGridFromEmail   string
	SendGridFromName    string
	WebURL              string
	UniDocMeteredKey    string
	Redis               RedisConfig
}

type GeneralConfig struct {
	Logger LoggerConfig
	Router RouterConfig
	Email  EmailConfig
}

type LoggerConfig struct {
	Level string
}

type RouterConfig struct {
	Port int
	CORS bool
}

type EmailConfig struct {
	Enable bool
}

type SQLConfig struct {
	Write DatabaseConfig
}

type RedisConfig struct {
	Write RedisCacheConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type ApplicationMessages struct {
	DuplicateUser           MessageFormat
	UserNotFound            MessageFormat
	InvalidLoginCredentials MessageFormat
	InternalServerError     MessageFormat
	BadRequestError         MessageFormat
	ValidationError         MessageFormat
	EmptyTokenError         MessageFormat
	InvalidTokenError       MessageFormat
	ExpireTokenError        MessageFormat
	UserIdEmptyError        MessageFormat
	UserNotVerified         MessageFormat
	UserAlreadyVerified     MessageFormat
	OtpUnMatchError         MessageFormat
}

type MessageFormat struct {
	Key     string
	Message string
}

type RedisCacheConfig struct {
	Host     string
	Port     string
	Password string
	Database int
}

// Initialize the global Viper instance
var v *viper.Viper
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "An application using Cobra and Viper",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := LoadConfig()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
		// fmt.Printf("Loaded config: %+v\n", config)
	},
}

func init() {
	v = viper.New()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is ./config.yaml)")
	v.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	// Adding a new flag -l for logging level
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "Set the logging level")
	v.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}

// initConfig initializes the configuration
func initConfig() {
	configPath := viper.GetString("config") // Get config path from Viper

	if configPath == "" {
		configPath = "././config" // Default path if not provided
	}

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("ABHI")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

// LoadConfig loads the configuration into the provided config struct
func LoadConfig() (*Config, error) {
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
