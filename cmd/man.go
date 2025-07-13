package main

import (
	"log"

	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1"

	"sigs.k8s.io/kubebuilder/v4/pkg/cli"
	cfgv3 "sigs.k8s.io/kubebuilder/v4/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin"
	kustomizev2 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/common/kustomize/v2"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	c := GetPluginsCLI()
	return c.Run()
}

func GetPluginsCLI() *cli.CLI {
	ansibleBundle, _ := plugin.NewBundleWithOptions(
		plugin.WithName(ansible.Plugin{}.Name()),
		plugin.WithVersion(ansible.Plugin{}.Version()),
		plugin.WithPlugins(
			kustomizev2.Plugin{},
			ansible.Plugin{},
		),
	)

	c, err := cli.New(
		cli.WithCommandName("ansible-cli"),
		cli.WithVersion(makeVersionString()),
		cli.WithPlugins(
			ansibleBundle,
		),
		cli.WithDefaultPlugins(cfgv3.Version, ansibleBundle),
		cli.WithDefaultProjectVersion(cfgv3.Version),
		cli.WithCompletion(),
	)

	if err != nil {
		log.Fatal(err)
	}

	return c
}

// TODO: @chiragkyal FIXME
func makeVersionString() string {
	// return fmt.Sprintf("ansible-cli version: %q, commit: %q, kubernetes version: %q, go version: %q, GOOS: %q, GOARCH: %q",
	// 	ver.GitVersion, ver.GitCommit, ver.KubernetesVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	return "0.0.0"
}
