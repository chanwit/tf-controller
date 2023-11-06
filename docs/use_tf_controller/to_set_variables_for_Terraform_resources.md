# Use TF-controller to set variables for Terraform resources

~> **BREAKING CHANGE**: This is a breaking change of the `v1alpha1` API.

Users who are upgrading from TF-controller <= 0.7.0 require updating `varsFrom`,
from a single object:

```yaml hl_lines="2"
  varsFrom:
    kind: ConfigMap
    name: cluster-config
```

to be an array of object, like this:

```yaml hl_lines="2"
  varsFrom:
  - kind: ConfigMap
    name: cluster-config
```

## `vars` and `varsFrom`

You can pass variables to Terraform using the `vars` and `varsFrom` fields.

Inline variables can be set using `vars`. The `varsFrom` field accepts a list of ConfigMaps / Secrets.
You may use the `varsKeys` property of `varsFrom` to select specific keys from the input or omit this field
to select all keys from the input source.

Values from `varsFrom` are always read first and then values from `vars` are read. If the same key is used in `varsFrom`
as in `vars`, the value from `varsFrom` will be overwritten. This mechanism allows us to have the global variables set 
once in one of the ConfigMaps or Secrets, and if necessary we can overwrite the value directly in the Terraform object.

```yaml hl_lines="15-20 22-28"
apiVersion: infra.contrib.fluxcd.io/v1alpha2
kind: Terraform
metadata:
  name: helloworld
  namespace: flux-system
spec:
  approvePlan: auto
  interval: 1m
  path: ./
  sourceRef:
    kind: GitRepository
    name: helloworld
    namespace: flux-system
  vars:
  - name: region
    value: us-east-1
  - name: env
    value: dev
  - name: instanceType
    value: t3-small
  varsFrom:
  - kind: ConfigMap
    name: cluster-config
    varsKeys:
    - nodeCount
    - instanceType
  - kind: Secret
    name: cluster-creds
```

## Variable value as HCL

The `vars` field supports HCL string, number, bool, object and list types. For example, the following variable can be populated using the accompanying Terraform spec:

```hcl hl_lines="3-6"
variable "cluster_spec" {
  type = object({
      region     = string
      env        = string
      node_count = number
      public     = bool
  })
}
```

```yaml hl_lines="17-20"
apiVersion: infra.contrib.fluxcd.io/v1alpha2
kind: Terraform
metadata:
  name: helloworld
  namespace: flux-system
spec:
  approvePlan: auto
  interval: 1m
  path: ./
  sourceRef:
    kind: GitRepository
    name: helloworld
    namespace: flux-system
  vars:
  - name: cluster_spec
    value:
      region: us-east-1
      env: dev
      node_count: 10
      public: false
```

## Rename variables in varsFrom

To rename a variable, you can use the varsKeys key within the varsFrom field. 
Here's the basic structure:

```yaml hl_lines="5"
spec:
  varsFrom:
  - kind: Secret
    name: <secret_name>
    varsKeys:
    - <original_variable_name>:<new_variable_name>
```
`original_variable_name` corresponds to the initial name of the variable in the referenced secret,
while `new_variable_name` represents the alias you want to use within the Terraform code.

Consider this example below, where we rename `nodeCount` to `node_count` 
and `instanceType` to `instance_type`:

```yaml hl_lines="18-19"
apiVersion: infra.contrib.fluxcd.io/v1alpha2
kind: Terraform
metadata:
  name: helloworld
  namespace: flux-system
spec:
  approvePlan: auto
  interval: 1m
  path: ./
  sourceRef:
    kind: GitRepository
    name: helloworld
    namespace: flux-system
  varsFrom:
  - kind: Secret
    name: cluster-config
    varsKeys:
    - nodeCount:node_count
    - instanceType:instance_type
```

## Rename output variables

See [Rename outputs](to_provision_resources_and_obtain_outputs.md#rename-outputs) for more details.

## Rename input secrets

See [Rename input secrets](with_the_ready_to_use_AWS_package.md#rename-input-secrets) for more details.
