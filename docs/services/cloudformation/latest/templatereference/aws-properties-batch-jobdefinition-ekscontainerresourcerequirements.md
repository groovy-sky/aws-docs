This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksContainerResourceRequirements

The type and amount of resources to assign to a container. The supported resources include
`memory`, `cpu`, and `nvidia.com/gpu`. For more information,
see [Resource management for pods and containers](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers) in the _Kubernetes_
_documentation_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Limits" : {Key: Value, ...},
  "Requests" : {Key: Value, ...}
}

```

### YAML

```yaml

  Limits:
    Key: Value
  Requests:
    Key: Value

```

## Properties

`Limits`

The type and quantity of the resources to reserve for the container. The values vary based
on the `name` that's specified. Resources can be requested using either the
`limits` or the `requests` objects.

memory

The memory hard limit (in MiB) for the container, using whole integers, with a "Mi"
suffix. If your container attempts to exceed the memory specified, the container is
terminated. You must specify at least 4 MiB of memory for a job. `memory` can be
specified in `limits`, `requests`, or both. If `memory` is
specified in both places, then the value that's specified in `limits` must be equal
to the value that's specified in `requests`.

###### Note

To maximize your resource utilization, provide your jobs with as much memory as possible
for the specific instance type that you are using. To learn how, see [Memory management](../../../batch/latest/userguide/memory-management.md) in the
_AWS Batch User Guide_.

cpu

The number of CPUs that's reserved for the container. Values must be an even multiple of
`0.25`. `cpu` can be specified in `limits`,
`requests`, or both. If `cpu` is specified in both places, then the
value that's specified in `limits` must be at least as large as the value that's
specified in `requests`.

nvidia.com/gpu

The number of GPUs that's reserved for the container. Values must be a whole integer.
`memory` can be specified in `limits`, `requests`, or both.
If `memory` is specified in both places, then the value that's specified in
`limits` must be equal to the value that's specified in
`requests`.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Requests`

The type and quantity of the resources to request for the container. The values vary based
on the `name` that's specified. Resources can be requested by using either the
`limits` or the `requests` objects.

memory

The memory hard limit (in MiB) for the container, using whole integers, with a "Mi"
suffix. If your container attempts to exceed the memory specified, the container is
terminated. You must specify at least 4 MiB of memory for a job. `memory` can be
specified in `limits`, `requests`, or both. If `memory` is
specified in both, then the value that's specified in `limits` must be equal to the
value that's specified in `requests`.

###### Note

If you're trying to maximize your resource utilization by providing your jobs as much
memory as possible for a particular instance type, see [Memory management](../../../batch/latest/userguide/memory-management.md) in the
_AWS Batch User Guide_.

cpu

The number of CPUs that are reserved for the container. Values must be an even multiple
of `0.25`. `cpu` can be specified in `limits`,
`requests`, or both. If `cpu` is specified in both, then the value
that's specified in `limits` must be at least as large as the value that's
specified in `requests`.

nvidia.com/gpu

The number of GPUs that are reserved for the container. Values must be a whole integer.
`nvidia.com/gpu` can be specified in `limits`, `requests`,
or both. If `nvidia.com/gpu` is specified in both, then the value that's specified
in `limits` must be equal to the value that's specified in
`requests`.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksContainerEnvironmentVariable

EksContainerSecurityContext

All content copied from https://docs.aws.amazon.com/.
