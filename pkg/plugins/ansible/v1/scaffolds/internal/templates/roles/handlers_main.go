// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package roles

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"

	"github.com/chiragkyal/ansible-plugin/pkg/plugins/ansible/v1/constants"
)

var _ machinery.Template = &HandlersMain{}

type HandlersMain struct {
	machinery.TemplateMixin
	machinery.ResourceMixin
}

// SetTemplateDefaults implements machinery.Template
func (f *HandlersMain) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(constants.RolesDir, "%[kind]", "handlers", "main.yml")
		f.Path = f.Resource.Replacer().Replace(f.Path)
	}

	f.TemplateBody = handlersMainAnsibleTmpl
	return nil
}

const handlersMainAnsibleTmpl = `---
# handlers file for {{ .Resource.Kind }}
`
