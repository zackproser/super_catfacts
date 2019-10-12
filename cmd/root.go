package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AppName is the common human legible name for this service
var AppName = "Super Catfacts Service"

var log = logrus.New()

// Config is the central configuration object
var Config Configuration

var catfacts, responses []string

var loglevel string
var verbose bool

var rootCmd = &cobra.Command{
	Use:              "super-catfacts",
	Short:            "Catfacts prank service",
	Long:             "Super catfacts is a full featured catfacts SMS and phone pranking service capable of running multiple simultaneous attacks.",
	PersistentPreRun: persistentPreRun,
}

// Execute - application entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
		return
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// TODO remove after dev
	var cfgLogLevel = logrus.DebugLevel

	if verbose {
		cfgLogLevel = logrus.DebugLevel
	}

	switch strings.ToLower(viper.GetString("loglevel")) {
	case "trace":
		cfgLogLevel = logrus.TraceLevel
	case "debug":
		cfgLogLevel = logrus.DebugLevel
	case "info":
		break
	case "warn":
		cfgLogLevel = logrus.WarnLevel
	case "error":
		cfgLogLevel = logrus.ErrorLevel
	case "fatal":
		cfgLogLevel = logrus.FatalLevel
	case "panic":
		cfgLogLevel = logrus.PanicLevel
	}

	log.SetLevel(cfgLogLevel)

	// Look for a config file in the working directory
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())

		marshalErr := viper.Unmarshal(&Config)
		if marshalErr != nil {
			log.Fatal("Unable to decode config into struct")
		}

		log.Debug("Read config file successfully")
		log.Debug("Admins are %v", Config.Server.Admins)
	} else {
		log.Debug("Error reading config file: %v", err)
	}

	if verbose {
		log.Debug("Config is :%v", Config)
	}
}

// Perform validation and initialization on required arguments
func persistentPreRun(cmd *cobra.Command, args []string) {

	if Config.Twilio.APIKey == "" {
		log.Fatal("Twilio API Key is a required argument")
	}

	if Config.Twilio.SID == "" {
		log.Fatal("Twilio SID is a required argument")
	}

	if Config.Server.FQDN == "" {
		log.Fatal("The fully qualified domain name (FQDN) of your catfacts server is a required argument")
	}

	if Config.Twilio.Number == "" {
		log.Fatal("A valid Twilio FROM number is a required argument")
	}

	if Config.Twilio.MsgIntervalSeconds == 0 {
		Config.Twilio.MsgIntervalSeconds = 30
	}

	if Config.Server.Port == "" {
		Config.Server.Port = ":8080"
	} else {
		Config.Server.Port = ":" + Config.Server.Port
	}

	for i := 0; i < len(Config.Server.Admins); i++ {
		valid, formatted := validateNumber(Config.Server.Admins[i])
		if valid {
			Config.Server.Admins[i] = formatted
		}
		log.WithFields(logrus.Fields{
			"Raw":       Config.Server.Admins[i],
			"Valid":     valid,
			"Formatted": formatted,
			"Parsed":    Config.Server.Admins,
		}).Debug("Parsing authorized server administrators")
	}
}
