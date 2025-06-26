package utils

import "strings"

type gitUrl string
type Project string
type Language string

const (
	NodeJS Language = "js"
	Golang Language = "golang"
)

const (
	URLStartWithGolang gitUrl = "https://github.com/scrapeless-ai/actor-template-go.git"
	URLStartWithNode   gitUrl = "https://github.com/scrapeless-ai/actor-template-ts.git"
)

const (
	ProjectStartWithGolang Project = "start_with_golang"
	ProjectStartWithNode   Project = "start_with_node"
)

var (
	ProjectMap = map[Project]gitUrl{
		ProjectStartWithGolang: URLStartWithGolang,
		ProjectStartWithNode:   URLStartWithNode,
	}
)

var (
	DevLanguage = map[Project]Language{
		ProjectStartWithGolang: Golang,
		ProjectStartWithNode:   NodeJS,
	}
)

func GetProjects() string {
	var projects []string
	for project := range ProjectMap {
		projects = append(projects, string(project))
	}
	return strings.Join(projects, ", ")
}
