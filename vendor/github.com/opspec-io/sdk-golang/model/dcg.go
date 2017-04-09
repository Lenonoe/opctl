package model

// dynamic call graph; see https://en.wikipedia.org/wiki/Call_graph
type DCG struct {
	Container *DCGContainer `json:"container,omitempty"`
	Op        *DCGOp        `json:"op,omitempty"`
}

type DCGBaseCall struct {
	RootOpId string `json:"rootOpId"`
	PkgRef   string `json:"pkgRef"`
}

type DCGUsernamePasswordAuth struct {
	// will be interpolated
	Username string `yaml:"username"`
	// will be interpolated
	Password string `yaml:"password"`
}

type DCGContainer struct{}

type DCGContainerCall struct {
	*DCGBaseCall
	ContainerId string   `json:"containerId"`
	Cmd         []string `json:"cmd"`
	// format: containerPath => hostPath
	Dirs map[string]string `json:"dirs"`
	// format: name => value
	EnvVars map[string]string `json:"envVars"`
	// format: containerPath => hostPath
	Files map[string]string      `json:"files"`
	Image *DCGContainerCallImage `json:"image"`
	// format: containerSocket => hostSocket
	Sockets map[string]string `json:"sockets"`
	WorkDir string            `json:"workDir"`
	Name    string            `json:"name,omitempty"`
	Ports   map[string]string `json:"ports,omitempty"`
}

type DCGContainerCallImage struct {
	// will be interpolated
	Ref      string                   `yaml:"ref"`
	PullAuth *DCGUsernamePasswordAuth `yaml:"pullAuth,omitempty"`
}

type DCGOp struct{}

type DCGOpCall struct {
	*DCGBaseCall
	OpId        string `json:"opId"`
	ChildCallId string `json:"childCallId"`
}

type DCGOpCallPkg struct {
	// will be interpolated
	Ref      string                   `yaml:"ref"`
	PullAuth *DCGUsernamePasswordAuth `yaml:"pullAuth,omitempty"`
}