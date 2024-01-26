package initial

import "github.com/wangyi12358/ratelctl/pkg/array"

type Project struct {
	Name string
	Url  string
}

var projects = []Project{
	{
		Name: "React Admin",
		Url:  "swc-admin",
	},
	{
		Name: "Nest Admin",
		Url:  "nest-admin-example",
	},
	{
		Name: "Go Gin",
		Url:  "go-gin-example",
	},
	{
		Name: "Go-zero",
		Url:  "go-zero-micro",
	},
	{
		Name: "React H5",
		Url:  "react-h5-template",
	},
}

var names = array.Map(projects, func(project Project) string {
	return project.Name
})
