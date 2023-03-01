package cmd

import (
	"fmt"
	"litectl/cmd/api"
	"litectl/internal/die"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Root = &cobra.Command{
	Use:   "litectl",
	Short: "A tiny IONOS CLI that can do a lot",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	Root.AddCommand(api.Api)
}

func initConfig() {
	const configFileType = "yaml"
	viper.SetConfigType(configFileType)
	viper.SetConfigFile(".lite-cli." + configFileType)

	viper.AutomaticEnv()
	die.AllMust(
		viper.BindEnv("IONOS_URL"),
		viper.BindEnv("IONOS_TOKEN"),
		viper.BindEnv("IONOS_USERNAME"),
		viper.BindEnv("IONOS_PASSWORD"),
	)

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}
