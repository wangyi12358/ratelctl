package initial

import "github.com/AlecAivazis/survey/v2"

type QuestionConfig struct {
	Project string
	Name    string
}

var questions = []*survey.Question{
	{
		Name: "project",
		Prompt: &survey.Select{
			Message: "请选择项目模版",
			Options: names,
		},
	},
	{
		Name: "name",
		Prompt: &survey.Input{
			Message: "请输入项目名称",
		},
	},
}
