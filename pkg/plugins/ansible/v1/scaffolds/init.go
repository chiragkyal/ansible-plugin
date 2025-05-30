/*
Copyright 2019 The Kubernetes Authors.
Modifications copyright 2020 The Operator-SDK Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scaffolds

import (
	"sigs.k8s.io/kubebuilder/v4/pkg/config"
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugins"
	kustomizev2 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/common/kustomize/v2"

	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates"
	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates/config/rbac"
	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates/config/testing"
	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates/config/testing/pullpolicy"
	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates/molecule/mdefault"
	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates/molecule/mkind"
	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates/playbooks"
	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/scaffolds/internal/templates/roles"
	// "github.com/operator-framework/ansible-operator-plugins/internal/version"
)

const imageName = "controller:latest"

// @chiragkyal TODO:
const imageVersion = "v0.0.1"

// ansibleOperatorVersion is set to the version of ansible-operator at compile-time.
var ansibleOperatorVersion = imageVersion

var _ plugins.Scaffolder = &initScaffolder{}

type initScaffolder struct {
	fs machinery.Filesystem

	config config.Config
}

// NewInitScaffolder returns a new plugins.Scaffolder for project initialization operations
func NewInitScaffolder(config config.Config) plugins.Scaffolder {
	return &initScaffolder{
		config: config,
	}
}

// InjectFS implements plugins.Scaffolder
func (s *initScaffolder) InjectFS(fs machinery.Filesystem) {
	s.fs = fs
}

// Scaffold implements plugins.Scaffolder
func (s *initScaffolder) Scaffold() error {
	// Initialize the machinery.Scaffold that will write the files to disk
	scaffold := machinery.NewScaffold(s.fs,
		// NOTE: kubebuilder's default permissions are only for root users
		machinery.WithDirectoryPermissions(0755),
		machinery.WithFilePermissions(0644),
		machinery.WithConfig(s.config),
	)

	return scaffold.Execute(
		&templates.Dockerfile{AnsibleOperatorVersion: ansibleOperatorVersion},
		&templates.Makefile{
			Image:                  imageName,
			KustomizeVersion:       kustomizev2.KustomizeVersion,
			AnsibleOperatorVersion: ansibleOperatorVersion,
		},
		&templates.GitIgnore{},
		&templates.RequirementsYml{},
		&templates.Watches{},
		&rbac.ManagerRole{},
		&roles.Placeholder{},
		&playbooks.Placeholder{},
		&mdefault.Converge{},
		&mdefault.Create{},
		&mdefault.Destroy{},
		&mdefault.Kustomize{},
		&mdefault.Molecule{},
		&mdefault.Prepare{},
		&mdefault.Verify{},
		&mkind.Converge{},
		&mkind.Create{},
		&mkind.Destroy{},
		&mkind.Molecule{},
		&pullpolicy.AlwaysPullPatch{},
		&pullpolicy.IfNotPresentPullPatch{},
		&pullpolicy.NeverPullPatch{},
		&testing.DebugLogsPatch{},
		&testing.Kustomization{},
		&testing.ManagerImage{},
	)
}
