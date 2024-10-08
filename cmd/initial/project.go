package initial

import "github.com/wangyi12358/ratelctl/pkg/array"

type Project struct {
	Name string
	Url  string
}

var projects = []Project{
	{
		Name: "React, Admin , Ant Design, RESTful",
		Url:  "swc-admin",
	},
	{
		Name: "Nest.js, Monolithic, RESTful",
		Url:  "nest-admin-example",
	},
	{
		Name: "Nest.js, Monolithic, GraphQL",
		Url:  "nestjs-graphql-single",
	},
	{
		Name: "Go, Gin, RESTful",
		Url:  "go-gin-example",
	},
	{
		Name: "Go, Go-Zero, Microservice",
		Url:  "go-zero-micro",
	},
	{
		Name: "React, H5",
		Url:  "react-h5-template",
	},
}

var names = array.Map(projects, func(project Project) string {
	return project.Name
})
