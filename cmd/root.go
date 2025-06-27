package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/smash-hq/scrape-cli/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	versionFlag  bool
	Version      = "v1.0.3"
	template     utils.Project
	templateName string
	createFlag   bool
	runFlag      bool

	defaultActorName = "my-actor"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "scrape-cli",
	Short: "Command-line interface for managing Scrapeless actors",
	Long: `scrape-cli is a command-line tool for creating, building, and running Scrapeless actor projects.
It supports interactive project generation, template-based initialization, and quick local execution.`,

	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Printf("scrapeless version: %s\n", Version)
			return
		}
		if createFlag {
			interactiveCreateTemplate()
			return
		}
		if cmd.Flags().Changed("tmpl") {
			createTemplate()
			return
		}
		if runFlag {
			utils.AutoRunProject()
			return
		}
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false,
		"Print the version number of scrape-cli")

	rootCmd.PersistentFlags().StringVarP((*string)(&template), "tmpl", "t", "",
		"Specify the template type to generate the actor code template, this is required when creating via command line\nSupported values: "+utils.GetProjectsStr())

	rootCmd.PersistentFlags().StringVarP(&templateName, "name", "n", defaultActorName,
		"Set the folder name for the generated actor code template")

	rootCmd.PersistentFlags().BoolVarP(&createFlag, "create", "c", false,
		"Generate a new actor code template interactively")

	rootCmd.PersistentFlags().BoolVarP(&runFlag, "run", "r", false,
		"Build and run the current actor code immediately")
}

func createTemplate() {
	url, ok := utils.ProjectMap[template]
	language := utils.DevLanguage[template]
	if !ok {
		fmt.Printf("Could not find the selected template: %s. Support list: %s\n", template, utils.GetProjectsStr())
		return
	}
	utils.CreateTemplate(url, templateName, language)
}

func interactiveCreateTemplate() {
	// 交互选择模板
	prompt := promptui.Select{
		Label: "Select a template",
		Items: utils.GetProjects(),
	}
	_, templateResult, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}
	template = utils.Project(templateResult)

	// 输入项目名
	namePrompt := promptui.Prompt{
		Label:   "Project name",
		Default: defaultActorName,
	}
	projectNameResult, err := namePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}
	templateName = projectNameResult

	createTemplate()
}
