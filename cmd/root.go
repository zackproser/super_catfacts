package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AppName = "Super Catfacts Service"

var log = logrus.New()

// Config is the central configuration object
var Config Configuration

var cfgFile, twilioAPIKey, twilioSid, port, loglevel string
var msgIntervalSeconds int
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

	// Set config defaults
	viper.SetDefault("port", "8080")

	viper.SetDefault("loglevel", "debug")

	viper.SetDefault("interval", 30)

	rootCmd.PersistentFlags().StringVarP(&twilioAPIKey, "apikey", "k", "", "Twilio API Key (Required)")

	rootCmd.PersistentFlags().StringVarP(&twilioSid, "sid", "s", "", "Twilio SID (Required)")

	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "Port to listen on")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode assists with debugging by dumping setup and conifguration info")

	rootCmd.PersistentFlags().IntVarP(&msgIntervalSeconds, "interval", "i", 30, "Number of seconds to wait between attack messages")

	rootCmd.PersistentFlags().StringVarP(&loglevel, "loglevel", "l", "debug", "Log Level")

	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

	viper.BindPFlag("sid", rootCmd.PersistentFlags().Lookup("sid"))

	viper.BindPFlag("interval", rootCmd.PersistentFlags().Lookup("interval"))

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
	viper.SetEnvPrefix("super_catfacts")

	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())

		marshalErr := viper.Unmarshal(&Config)
		if marshalErr != nil {
			log.Fatal("Unable to decode config into struct %v", marshalErr)
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

	if Config.Twilio.APIKey != "" {
		twilioAPIKey = Config.Twilio.APIKey
	} else {
		twilioAPIKey = viper.GetString("apikey")
	}

	if twilioAPIKey == "" {
		log.Fatal("Twilio API Key is a required argument")
	}

	if Config.Twilio.SID != "" {
		twilioSid = Config.Twilio.SID
	} else {
		twilioSid = viper.GetString("sid")
	}

	if twilioSid == "" {
		log.Fatal("Twilio SID is a required argument")
	}

	if Config.Twilio.MsgIntervalSeconds == 0 {
		Config.Twilio.MsgIntervalSeconds = viper.GetInt("interval")
	}

	if Config.Server.Port != "" {
		port = Config.Server.Port
	} else {
		port = viper.GetString("port")
	}
	Config.Server.Port = ":" + port

}
