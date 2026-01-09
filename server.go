package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "caching-server",
	Short: "A simple caching server",
	Long:  `A simple caching server that allows you to store and retrieve data with ease.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Caching Server!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.caching-server.yaml)")
}

func initConfig() {
	// Configuration initialization logic can be added here
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".caching-server")
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CACHEPROXY")

	viper.SetDefault("precision", 2)
	viper.SetDefault("factor", 1.0)

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
