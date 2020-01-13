package entity

type Edge struct {
	ID       string
	FromNode string
	ToNode   string
}

type Node struct {
	ID     string
	Name   string
	Image  string
	Src    string
	Config map[string]string
}

type Workflow struct {
	Name       string
	Entrypoint string
	Nodes      []*Node
	Edges      []*Edge
}

type Entrypoint struct {
	ProtoFile string
	Image     string
	Src       string
	Config    map[string]interface{}
}

type Version struct {
	Name       string
	Entrypoint Entrypoint
	Workflows  []Workflow
	Status     string
}
