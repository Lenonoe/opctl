package docker

import (
	"fmt"
	"github.com/docker/docker/api/types/container"
	dockerClientPkg "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"sort"
	"strings"
)

//counterfeiter:generate -o internal/fakes/hostConfigFactory.go . hostConfigFactory
type hostConfigFactory interface {
	Construct(
		containerCallDirs map[string]string,
		containerCallFiles map[string]string,
		containerCallSockets map[string]string,
		portBindings nat.PortMap,
	) *container.HostConfig
}

func newHostConfigFactory(
	dockerClient dockerClientPkg.CommonAPIClient,
) (hostConfigFactory, error) {
	fspc, err := newFSPathConverter(dockerClient)
	if err != nil {
		return _hostConfigFactory{}, err
	}

	hcf := _hostConfigFactory{
		fsPathConverter: fspc,
	}
	return hcf, nil
}

type _hostConfigFactory struct {
	fsPathConverter fsPathConverter
}

func (hcf _hostConfigFactory) Construct(
	containerCallDirs map[string]string,
	containerCallFiles map[string]string,
	containerCallSockets map[string]string,
	portBindings nat.PortMap,
) *container.HostConfig {
	hostConfig := &container.HostConfig{
		PortBindings: portBindings,
		// support docker in docker
		// @TODO: reconsider; can we avoid this?
		// see for similar discussion: https://github.com/kubernetes/kubernetes/issues/391
		Privileged: true,
	}
	for containerFilePath, hostFilePath := range containerCallFiles {
		hostConfig.Binds = append(
			hostConfig.Binds,
			fmt.Sprintf(
				"%v:%v:cached",
				hcf.fsPathConverter.LocalToEngine(hostFilePath),
				containerFilePath,
			),
		)
	}
	for containerDirPath, hostDirPath := range containerCallDirs {
		hostConfig.Binds = append(
			hostConfig.Binds,
			fmt.Sprintf(
				"%v:%v:cached",
				hcf.fsPathConverter.LocalToEngine(hostDirPath),
				containerDirPath,
			),
		)
	}
	for containerSocketAddress, hostSocketAddress := range containerCallSockets {
		const unixSocketAddressDiscriminationChars = `/\`
		// note: this mechanism for determining the type of socket is naive; higher level of sophistication may be required
		if strings.ContainsAny(hostSocketAddress, unixSocketAddressDiscriminationChars) {
			hostConfig.Binds = append(
				hostConfig.Binds,
				fmt.Sprintf(
					"%v:%v",
					hcf.fsPathConverter.LocalToEngine(hostSocketAddress),
					containerSocketAddress,
				),
			)
		}
	}
	// sort binds to make order deterministic; useful for testing
	sort.Strings(hostConfig.Binds)

	return hostConfig
}
