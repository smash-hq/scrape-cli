package utils

import "strings"

type gitUrl string
type Project string

const (
	URLStartWithGolang gitUrl = "https://github.com/scrapeless-ai/actor-template-go.git"
	URLStartWithNode   gitUrl = "https://github.com/scrapeless-ai/sdk-node.git"
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

func GetProjects() string {
	var projects []string
	for project := range ProjectMap {
		projects = append(projects, string(project))
	}
	return strings.Join(projects, ", ")
}
