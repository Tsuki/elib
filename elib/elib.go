package main

import (
	"elib/log"
	"elib/utils"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "elib",
		Short: "jellyfin database",
		Long:  "jellyfin database",
	}

	version = "dev"
	date    = "unknown"
	cfgFile = ""
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(runCmd)
	//rootCmd.AddCommand(migrateCmd)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is elib.yaml)")
	runCmd.Flags().BoolVarP(&utils.Option.Debug, "debug", "d", false, "Debug")
	log.CheckFatal(viper.BindPFlag("options.debug", runCmd.Flags().Lookup("debug")))
	viper.SetDefault("options", utils.OptionsDefault)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	for _, value := range figure.NewFigure("elib", "speed", true).Slicify() {
		println(value)
	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		cfgFile, _ = os.Getwd()
		log.Infof("pwd: %s", cfgFile)
		viper.AddConfigPath(cfgFile)
		viper.SetConfigName("elib")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Infof("Creat default config file %s", cfgFile)
			log.CheckFatal(viper.WriteConfigAs(cfgFile + "/elib.yaml"))
		} else {
			log.Fatalf("Viper unexpected error %s", err.Error())
		}
	}
	log.Info("Using config file:", viper.ConfigFileUsed())
	log.CheckFatal(viper.Unmarshal(&struct{ *utils.Options }{&utils.Option}))
}

func Execute(version, date string) {
	rootCmd.Version = fmt.Sprintf("version %v, date %v", version, date)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Execute(version, date)
}
