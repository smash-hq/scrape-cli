package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
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
	var selected string
	var projectName string

	// 选择模板
	templates := utils.GetProjects()
	templatePrompt := &survey.Select{
		Message: "Select a template",
		Help:    "Select the project code template you want",
		Default: string(utils.ProjectStartWithGolang),
		Options: templates,
	}
	if err := survey.AskOne(templatePrompt, &selected); err != nil {
		fmt.Printf("❌ Template selection failed: %v\n", err)
		return
	}
	template = utils.Project(selected)

	// 输入项目名
	namePrompt := &survey.Input{
		Message: "Enter the name of the template",
		Default: defaultActorName,
		Help: "If it is blank, the default name will be used\n" +
			"1. Only lowercase English letters (a–z) are allowed.\n" +
			"2. Hyphens `-` and underscores `_` are permitted as separators.\n" +
			"3. No uppercase letters, numbers, spaces, special characters, or non-English characters (e.g. Chinese) are allowed.",
	}
	if err := survey.AskOne(namePrompt, &projectName); err != nil {
		fmt.Printf("❌ Project name input failed: %v\n", err)
		return
	}
	templateName = projectName

	createTemplate()
}
