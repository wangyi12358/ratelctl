package initial

import "github.com/wangyi12358/ratelctl/pkg/array"

type Project struct {
	Name string
	Url  string
}

var projects = []Project{
	{
		Name: "React Admin",
		Url:  "https://github.com/wangyi12358/swc-admin.git",
	},
	{
		Name: "Nest Admin",
		Url:  "https://github.com/wangyi12358/nest-admin-example.git",
	},
	{
		Name: "Go Gin",
		Url:  "https://github.com/wangyi12358/go-gin-example.git",
	},
	{
		Name: "Go-zero",
		Url:  "https://github.com/wangyi12358/go-zero-micro.git",
	},
}

var names = array.Map(projects, func(project Project) string {
	return project.Name
})
