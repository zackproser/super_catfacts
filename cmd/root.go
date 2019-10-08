package cmd

import (
	"strings"
	"super_catfacts/common"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logrus.New()
var cfgFile, twilioAPIKey, twilioSid, port, bindPort, logLevel string
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
	log.Debug("Executing")
	if err := rootCmd.Execute(); err != nil {
		log.Debug(err)
		return
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&twilioAPIKey, "apikey", "k", "", "Twilio API Key (Required)")

	rootCmd.PersistentFlags().StringVarP(&twilioSid, "sid", "s", "", "Twilio SID (Required)")

	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "Port to listen on")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose mode assists with debugging by dumping setup and conifguration info")

	rootCmd.PersistentFlags().IntVarP(&msgIntervalSeconds, "interval", "i", 30, "Number of seconds to wait between attack messages")

	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

	viper.BindPFlag("sid", rootCmd.PersistentFlags().Lookup("sid"))

	viper.BindPFlag("interval", rootCmd.PersistentFlags().Lookup("interval"))

}

func initConfig() {

	var cfgLogLevel = logrus.DebugLevel

	if verbose {
		cfgLogLevel = logrus.DebugLevel
	}

	switch strings.ToLower(logLevel) {
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

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("super_catfacts")

	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())

		marshalErr := viper.Unmarshal(&common.Configuration)
		if marshalErr != nil {
			log.Fatal("Unable to decode config into struct %v", marshalErr)
		}

		log.Debug("Read config file successfully")
		log.Debug("Admins are %v", common.Configuration.Server.Admins)
	} else {
		log.Debug("Error reading config file: %v", err)
	}

	if verbose {
		log.Debug("Configuration is :%v", common.Configuration)
	}
}

// Perform validation and initialization on required arguments
func persistentPreRun(cmd *cobra.Command, args []string) {

	if common.Configuration.Twilio.APIKey != "" {
		twilioAPIKey = common.Configuration.Twilio.APIKey
	} else {
		twilioAPIKey = viper.GetString("apikey")
	}
	if twilioAPIKey == "" {
		log.Fatal("Twilio API Key is a required argument")
	}

	if common.Configuration.Twilio.SID != "" {
		twilioSid = common.Configuration.Twilio.SID
	} else {
		twilioSid = viper.GetString("sid")
	}

	if twilioSid == "" {
		log.Fatal("Twilio SID is a required argument")
	}

	/*if common.Configuration.Twilio.MsgIntervalSeconds != nil {
		msgIntervalSeconds = common.Configuration.Twilio.MsgIntervalSeconds
	} else {
		msgIntervalSeconds = viper.GetInt("interval")
	} */

	if common.Configuration.Server.Port != "" {
		bindPort = common.Configuration.Server.Port
	} else {
		bindPort = viper.GetString("port")
	}

	if bindPort == "" {
		log.Debug("No port supplied. Defaulting to listing on ", common.DefaultPort)
		bindPort = common.DefaultPort
	}
	common.Configuration.Server.Port = ":" + bindPort

}
