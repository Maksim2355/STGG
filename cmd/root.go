/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "STGG",
	Short: "CLI для простой шаблонизации",
	Long: `Шаблонизатор с использованием стандартного языка шаблонов golang
	
	Более подробную документацию можно найти в репозитории проекта https://github.com/Maksim2355/STGG`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Установка шаблонизатора прошла успешно")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//nothing
}
