package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"stgg/cmd/printer"
	"stgg/crossplatform"
	"stgg/res"
)

var (
	// Used for flags.
	cfgFile     string
	cfgFileName = "stggCfg"
)
var rootCmd = &cobra.Command{
	Use:   "STGG",
	Short: "CLI для простой шаблонизации",
	Long: `Шаблонизатор с использованием стандартного языка шаблонов golang
	
	Более подробную документацию можно найти в репозитории проекта https://github.com/Maksim2355/STGG`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		printer.PrintErrorAndExit(err.Error())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/stggCfg.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home := res.HomeDir()
		newPathCfgFile := home + crossplatform.PATH_SEPARATOR + cfgFileName + ".yaml"
		if _, err := os.Stat(newPathCfgFile); errors.Is(err, os.ErrNotExist) {
			file, err := os.Create(newPathCfgFile)
			if err != nil {
				printer.PrintError("Ошибка создания файла конфигурации. Создайте файл сами")
			} else {
				defaultConfig := res.PrepareDefaultConfig()
				file.Write([]byte(defaultConfig))
			}
		}
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(cfgFileName)
	}

	viper.AutomaticEnv()
}
