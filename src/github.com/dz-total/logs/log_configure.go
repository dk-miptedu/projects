package logs

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Log *logrus.Logger

func InitializeLogging(path string) {

	// Set up Viper to read the config file
	viper.SetConfigName("logsetup")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path) // or wherever your config file is located

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}

	// Configure logrus according to the config file
	configureLogrus()
}

func configureLogrus() {
	Log = logrus.New()

	// Set the log level according to the config file
	levelString := viper.GetString("logger.logLevel")
	level, err := logrus.ParseLevel(levelString)
	if err != nil {
		logrus.Fatalf("Invalid log level: %s", levelString)
	}
	Log.SetLevel(level)
	Log.SetOutput(os.Stdout)
	logFile := viper.GetString("logger.logFile")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}
	Log.SetOutput(file)

	// Optionally, set the formatter as JSON or Text
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
