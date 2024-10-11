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
		Url:  "nestjs-graphql-monolith",
	},
	{
		Name: "Nest.js, Monolithic, GraphQL",
		Url:  "nestjs-restful-monolith",
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
		Name: "Go, gRPC, Microservice",
		Url:  "go-grpc-microservices",
	},
	{
		Name: "React, H5",
		Url:  "react-h5-template",
	},
}

var names = array.Map(projects, func(project Project) string {
	return project.Name
})
