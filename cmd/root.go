package cmd

import (
	"fmt"
	"github.com/smash-hq/scrape-cli/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	versionFlag bool
	Version     = "v1.0.0" // 自定义版本号
	template    utils.Project
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scrape-cli",
	Short: "Scrapeless operation from the command line",
	Long:  `Scrapeless operation from the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Printf("scrapeless version: %s\n", Version)
			return
		}
		// 2. 处理 template 模式
		if template != "" {
			createTemplate()
			return
		}
		// 默认显示帮助
		_ = cmd.Help()
	},
}

func createTemplate() {
	switch template {
	case utils.ProjectStartWithGolang:
		url := utils.ProjectMap[utils.ProjectStartWithGolang]
		utils.CreateGoTemplate(url)
	case utils.ProjectStartWithNode:
		utils.CreateNodeTemplate()
	default:
		fmt.Printf("Could not find the selected template: %s in the list of templates. support list: %s.\n", template, utils.GetProjects())
	}
	return
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Print the version number of scrapeless")
	rootCmd.PersistentFlags().StringVarP((*string)(&template), "tmpl", "t", "", "Generate template folder (e.g. --template go)")

}
