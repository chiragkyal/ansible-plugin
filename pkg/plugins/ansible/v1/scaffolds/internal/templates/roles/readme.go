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

const ReadmePath = "README.md"

var _ machinery.Template = &Readme{}

type Readme struct {
	machinery.TemplateMixin
	machinery.ResourceMixin
}

func (f *Readme) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join(constants.RolesDir, "%[kind]", ReadmePath)
	}
	f.Path = f.Resource.Replacer().Replace(f.Path)

	f.TemplateBody = readmeAnsibleTmpl
	return nil
}

const readmeAnsibleTmpl = `Role Name
=========

A brief description of the role goes here.

Requirements
------------

Any pre-requisites that may not be covered by Ansible itself or the role should be mentioned here. For instance,
if the role uses the EC2 module, it may be a good idea to mention in this section that the boto package is required.

Role Variables
--------------

A description of the settable variables for this role should go here, including any variables that are in 
defaults/main.yml, vars/main.yml, and any variables that can/should be set via parameters to the role. Any variables 
that are read from other roles and/or the global scope (ie. hostvars, group vars, etc.) should be mentioned here as well

Dependencies
------------

A list of other roles hosted on Galaxy should go here, plus any details in regards to parameters that may need to be set
for other roles, or variables that are used from other roles.

Example Playbook
----------------

Including an example of how to use your role (for instance, with variables passed in as parameters) is always nice for
users too:

    - hosts: servers
      roles:
         - { role: username.rolename, x: 42 }

License
-------

BSD

Author Information
------------------

An optional section for the role authors to include contact information, or a website (HTML is not allowed).
`
