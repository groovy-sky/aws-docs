This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition LinuxParameters

The Linux-specific options that are applied to the container, such as Linux [KernelCapabilities](../../../../reference/amazonecs/latest/apireference/api-kernelcapabilities.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Capabilities" : KernelCapabilities,
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

  Capabilities:
    KernelCapabilities
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

`Capabilities`

The Linux capabilities for the container that are added to or dropped from the default
configuration provided by Docker.

###### Note

For tasks that use the Fargate launch type, `capabilities` is supported
for all platform versions but the `add` parameter is only supported if
using platform version 1.4.0 or later.

_Required_: No

_Type_: [KernelCapabilities](aws-properties-ecs-taskdefinition-kernelcapabilities.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Devices`

Any host devices to expose to the container. This parameter maps to
`Devices` in the docker container create command and the
`--device` option to docker run.

###### Note

If you're using tasks that use the Fargate launch type, the `devices`
parameter isn't supported.

_Required_: No

_Type_: Array of [Device](aws-properties-ecs-taskdefinition-device.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InitProcessEnabled`

Run an `init` process inside the container that forwards signals and reaps
processes. This parameter maps to the `--init` option to docker run. This
parameter requires version 1.25 of the Docker Remote API or greater on your container
instance. To check the Docker Remote API version on your container instance, log in to
your container instance and run the following command: `sudo docker version
				--format '{{.Server.APIVersion}}'`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxSwap`

The total amount of swap memory (in MiB) a container can use. This parameter will be
translated to the `--memory-swap` option to docker run where the value would
be the sum of the container memory plus the `maxSwap` value.

If a `maxSwap` value of `0` is specified, the container will not
use swap. Accepted values are `0` or any positive integer. If the
`maxSwap` parameter is omitted, the container will use the swap
configuration for the container instance it is running on. A `maxSwap` value
must be set for the `swappiness` parameter to be used.

###### Note

If you're using tasks that use the Fargate launch type, the `maxSwap`
parameter isn't supported.

If you're using tasks on Amazon Linux 2023 the `swappiness` parameter
isn't supported.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SharedMemorySize`

The value for the size (in MiB) of the `/dev/shm` volume. This parameter
maps to the `--shm-size` option to docker run.

###### Note

If you are using tasks that use the Fargate launch type, the
`sharedMemorySize` parameter is not supported.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Swappiness`

This allows you to tune a container's memory swappiness behavior. A
`swappiness` value of `0` will cause swapping to not happen
unless absolutely necessary. A `swappiness` value of `100` will
cause pages to be swapped very aggressively. Accepted values are whole numbers between
`0` and `100`. If the `swappiness` parameter is not
specified, a default value of `60` is used. If a value is not specified for
`maxSwap` then this parameter is ignored. This parameter maps to the
`--memory-swappiness` option to docker run.

###### Note

If you're using tasks that use the Fargate launch type, the
`swappiness` parameter isn't supported.

If you're using tasks on Amazon Linux 2023 the `swappiness` parameter
isn't supported.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tmpfs`

The container path, mount options, and size (in MiB) of the tmpfs mount. This
parameter maps to the `--tmpfs` option to docker run.

###### Note

If you're using tasks that use the Fargate launch type, the `tmpfs`
parameter isn't supported.

_Required_: No

_Type_: [Array](aws-properties-ecs-taskdefinition-tmpfs.md) of [Tmpfs](aws-properties-ecs-taskdefinition-tmpfs.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyValuePair

LogConfiguration

All content copied from https://docs.aws.amazon.com/.
