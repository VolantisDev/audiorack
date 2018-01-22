package config

import (
	"flag"
	"path/filepath"
	"github.com/spf13/viper"
	"github.com/shibukawa/configdir"
	"github.com/juju/errors"
	"github.com/spf13/pflag"
	"time"
)

func init() {
	appData := configdir.New("VolantisDev", "VstGuide")
	configPath := filepath.Join(appData.LocalPath, "Config")
	databasePath := filepath.Join(appData.LocalPath, "Database")

	setDefaults(configPath, databasePath)
	loadConfig(configPath)
	bindEnvironmentVars()
	bindFlags(configPath, databasePath)
}

func bindFlags(configPath string, databasePath string) {
	flag.String("configPath", configPath, "The path to the folder containing the configuration file(s).")
	flag.String("databasePath", databasePath, "The path to the folder containing the database file(s).")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func bindEnvironmentVars() {
	viper.SetEnvPrefix("VstGuide")

	viper.BindEnv("configPath")
	viper.BindEnv("databasePath")
	viper.BindEnv("vstPaths")
	viper.BindEnv("samplePaths")
	viper.BindEnv("kontaktLibraries")
}

func loadConfig(configPath string) {
	viper.SetConfigName("VstGuide")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.Annotate(err, "Fatal error reading config file!"))
	}
}

func setDefaults(configPath string, databasePath string) {
	// Internal config
	viper.SetDefault("configPath", configPath)
	viper.SetDefault("databasePath", databasePath)

	// Config for the user library
	viper.SetDefault("vstPaths", defaultVstPaths())
	viper.SetDefault("samplePaths", defaultSamplePaths())
	viper.SetDefault("kontaktLibraryPaths", defaultKontaktLibraryPaths())
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func GetDuration(key string) time.Duration{
	return viper.GetDuration(key)
}

func IsSet(key string) bool {
	return viper.IsSet(key)
}
