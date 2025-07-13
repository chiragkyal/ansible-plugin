# Ansible Plugin for Kubebuilder



## Installation

```bash
# Clone the repository
git clone https://github.com/chiragkyal/ansible-plugin.git
cd ansible-plugin

# Build the binary
make build

# Install to $GOPATH/bin (optional)
make install
```


## Quick Start

### 1. Generate a Sample Operator

The fastest way to get started is to generate a sample operator:

```bash
make generate-sample
```

This will create a `memcached-operator` directory with a complete Ansible operator project.

### 2. Manual Setup

If you prefer to set up manually:

```bash
# Build the CLI tool
make build

# Create a new operator project
mkdir my-operator
cd my-operator

# Initialize the operator
../bin/ansible-cli init --domain example.com

# Create an API
../bin/ansible-cli create api --group cache --version v1alpha1 --kind Memcached --generate-role
```
