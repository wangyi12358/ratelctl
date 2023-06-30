package initial

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/wangyi12358/ratelctl/pkg/array"
	"os"
	"os/exec"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "快速生成项目模板",
	Run:   InitCmd,
}

func InitCmd(_ *cobra.Command, _ []string) {
	var config QuestionConfig
	if err := survey.Ask(questions, &config); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	selected := array.Find(projects, func(p Project) bool {
		return p.Name == config.Project
	})
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Cloning project template"
	s.Start()
	if err := exec.Command("git", "clone", selected.Url, config.Name).Run(); err != nil {
		fmt.Println(err.Error())
		s.Stop()
		os.Exit(0)
	}
	s.Stop()
	green := color.FgGreen.Render
	fmt.Println(green(fmt.Sprintf("Successfully initialized %s project", config.Project)))
}
