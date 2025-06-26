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
	Short: "Scrapeless operation from the command line",
	Long:  `Scrapeless operation from the command line`,
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
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Print the version number of scrape-cli")
	rootCmd.PersistentFlags().StringVarP((*string)(&template), "tmpl", "t", "", "Generate code template (e.g. --tmpl start_with_golang), "+
		"When creating with command line, the parameter can not be empty")
	rootCmd.PersistentFlags().StringVarP(&templateName, "name", "n", defaultActorName, "Specify the folder name for the generated template")
	rootCmd.PersistentFlags().BoolVarP(&createFlag, "create", "c", false, "Generate code template by interactively")
	rootCmd.PersistentFlags().BoolVarP(&runFlag, "run", "r", false, "Quickly launch your actor")
}

func createTemplate() {
	url, ok := utils.ProjectMap[template]
	if !ok {
		fmt.Printf("Could not find the selected template: %s. Support list: %s\n", template, utils.GetProjects())
		return
	}
	utils.CreateTemplate(url, templateName)
	fmt.Printf("Project '%s' created using '%s' template.\n", templateName, template)
}

func interactiveCreateTemplate() {
	// 交互选择模板
	prompt := promptui.Select{
		Label: "Select a template",
		Items: []string{string(utils.ProjectStartWithGolang), string(utils.ProjectStartWithNode)},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}
	template = utils.Project(result)

	// 输入项目名
	namePrompt := promptui.Prompt{
		Label:   "Project name",
		Default: defaultActorName,
	}
	result, err = namePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}
	templateName = result

	createTemplate()
}
