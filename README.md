# Ansible Plugin For Scaffolding Ansible Operators

```
make build

mkdir memcached-operator
cd memcached-operator

../ansible-cli init --domain example.com

../ansible-cli create api --group cache --version v1alpha1 --kind Memcached --generate-role
```