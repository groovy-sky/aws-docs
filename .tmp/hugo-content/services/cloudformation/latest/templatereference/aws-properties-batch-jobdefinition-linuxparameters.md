This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition LinuxParameters

Linux-specific modifications that are applied to the container, such as details for device
mappings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Devices" : [ Device, ... ],
  "InitProcessEnabled" : Boolean,
  "MaxSwap" : Integer,
  "SharedMemorySize" : Integer,
  "Swappiness" : Integer,
  "Tmpfs" : [ Tmpfs, ... ]
}

```

### YAML

```yaml

  Devices:
    - Device
  InitProcessEnabled: Boolean
  MaxSwap: Integer
  SharedMemorySize: Integer
  Swappiness: Integer
  Tmpfs:
    - Tmpfs

```

## Properties

`Devices`

Any of the host devices to expose to the container. This parameter maps to
`Devices` in the [Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23)
and the `--device` option to [docker\
run](https://docs.docker.com/engine/reference/run).

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't
provide it for these jobs.

_Required_: No

_Type_: Array of [Device](aws-properties-batch-jobdefinition-device.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InitProcessEnabled`

If true, run an `init` process inside the container that forwards signals and
reaps processes. This parameter maps to the `--init` option to [docker run](https://docs.docker.com/engine/reference/run). This parameter requires version 1.25 of the Docker Remote API or greater on your
container instance. To check the Docker Remote API version on your container instance, log in to your
container instance and run the following command: `sudo docker version | grep "Server API version"`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSwap`

The total amount of swap memory (in MiB) a container can use. This parameter is translated
to the `--memory-swap` option to [docker\
run](https://docs.docker.com/engine/reference/run) where the value is the sum of the container memory plus the `maxSwap`
value. For more information, see [`--memory-swap` details](https://docs.docker.com/config/containers/resource_constraints) in the Docker documentation.

If a `maxSwap` value of `0` is specified, the container doesn't use
swap. Accepted values are `0` or any positive integer. If the `maxSwap`
parameter is omitted, the container doesn't use the swap configuration for the container instance
on which it runs. A `maxSwap` value must be set for the `swappiness`
parameter to be used.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't
provide it for these jobs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharedMemorySize`

The value for the size (in MiB) of the `/dev/shm` volume. This parameter maps to
the `--shm-size` option to [docker\
run](https://docs.docker.com/engine/reference/run).

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't
provide it for these jobs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Swappiness`

You can use this parameter to tune a container's memory swappiness behavior. A
`swappiness` value of `0` causes swapping to not occur unless absolutely
necessary. A `swappiness` value of `100` causes pages to be swapped
aggressively. Valid values are whole numbers between `0` and `100`. If the
`swappiness` parameter isn't specified, a default value of `60` is used.
If a value isn't specified for `maxSwap`, then this parameter is ignored. If
`maxSwap` is set to 0, the container doesn't use swap. This parameter maps to the
`--memory-swappiness` option to [docker\
run](https://docs.docker.com/engine/reference/run).

Consider the following when you use a per-container swap configuration.

- Swap space must be enabled and allocated on the container instance for the containers to
use.

###### Note

By default, the Amazon ECS optimized AMIs don't have swap enabled. You must enable swap on the
instance to use this feature. For more information, see [Instance store swap\
volumes](../../../ec2/latest/userguide/instance-store-swap-volumes.md) in the _Amazon EC2 User Guide for Linux Instances_ or [How do I\
allocate memory to work as swap space in an Amazon EC2 instance by using a swap\
file?](https://aws.amazon.com/premiumsupport/knowledge-center/ec2-memory-swap-file)

- The swap space parameters are only supported for job definitions using EC2
resources.

- If the `maxSwap` and `swappiness` parameters are omitted from a job
definition, each container has a default `swappiness` value of 60. Moreover, the
total swap usage is limited to two times the memory reservation of the container.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't
provide it for these jobs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tmpfs`

The container path, mount options, and size (in MiB) of the `tmpfs` mount. This
parameter maps to the `--tmpfs` option to [docker\
run](https://docs.docker.com/engine/reference/run).

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't
provide this parameter for this resource type.

_Required_: No

_Type_: [Array](aws-properties-batch-jobdefinition-tmpfs.md) of [Tmpfs](aws-properties-batch-jobdefinition-tmpfs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobTimeout

LogConfiguration

All content copied from https://docs.aws.amazon.com/.
