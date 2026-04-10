This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition TaskContainerProperties

Container properties are used for Amazon ECS-based job definitions. These properties to describe
the container that's launched as part of a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "DependsOn" : [ TaskContainerDependency, ... ],
  "Environment" : [ Environment, ... ],
  "Essential" : Boolean,
  "FirelensConfiguration" : FirelensConfiguration,
  "Image" : String,
  "LinuxParameters" : LinuxParameters,
  "LogConfiguration" : LogConfiguration,
  "MountPoints" : [ MountPoint, ... ],
  "Name" : String,
  "Privileged" : Boolean,
  "ReadonlyRootFilesystem" : Boolean,
  "RepositoryCredentials" : RepositoryCredentials,
  "ResourceRequirements" : [ ResourceRequirement, ... ],
  "Secrets" : [ Secret, ... ],
  "Ulimits" : [ Ulimit, ... ],
  "User" : String
}

```

### YAML

```yaml

  Command:
    - String
  DependsOn:
    - TaskContainerDependency
  Environment:
    - Environment
  Essential: Boolean
  FirelensConfiguration:
    FirelensConfiguration
  Image: String
  LinuxParameters:
    LinuxParameters
  LogConfiguration:
    LogConfiguration
  MountPoints:
    - MountPoint
  Name: String
  Privileged: Boolean
  ReadonlyRootFilesystem: Boolean
  RepositoryCredentials:
    RepositoryCredentials
  ResourceRequirements:
    - ResourceRequirement
  Secrets:
    - Secret
  Ulimits:
    - Ulimit
  User: String

```

## Properties

`Command`

The command that's passed to the container. This parameter maps to `Cmd` in the
[Create a\
container](https://docs.docker.com/engine/api/v1.23) section of the [Docker\
Remote API](https://docs.docker.com/engine/api/v1.23) and the `COMMAND` parameter to [docker run](https://docs.docker.com/engine/reference/run). For more information,
see [Dockerfile reference:\
CMD](https://docs.docker.com/engine/reference/builder).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DependsOn`

A list of containers that this container depends on.

_Required_: No

_Type_: Array of [TaskContainerDependency](aws-properties-batch-jobdefinition-taskcontainerdependency.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables to pass to a container. This parameter maps to Env in the [Create a container](https://docs.docker.com/engine/api/v1.23)
section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23)
and the `--env` parameter to [docker run](https://docs.docker.com/engine/reference/run).

###### Important

We don't recommend using plaintext environment variables for sensitive information, such as
credential data.

###### Note

Environment variables cannot start with `AWS_BATCH`. This naming convention is
reserved for variables that AWS Batch sets.

_Required_: No

_Type_: [Array](aws-properties-batch-jobdefinition-environment.md) of [Environment](aws-properties-batch-jobdefinition-environment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Essential`

If the essential parameter of a container is marked as `true`, and that container
fails or stops for any reason, all other containers that are part of the task are stopped. If the
`essential` parameter of a container is marked as false, its failure doesn't affect
the rest of the containers in a task. If this parameter is omitted, a container is assumed to be
essential.

All jobs must have at least one essential container. If you have an application that's
composed of multiple containers, group containers that are used for a common purpose into
components, and separate the different components into multiple task definitions. For more
information, see [Application\
Architecture](../../../amazonecs/latest/developerguide/application-architecture.md) in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirelensConfiguration`

The FireLens configuration for the container. This is used to specify and configure a
log router for container logs. For more information, see [Custom log](../../../amazonecs/latest/developerguide/using-firelens.md) routing in the _Amazon Elastic Container Service Developer_
_Guide_.

_Required_: No

_Type_: [FirelensConfiguration](aws-properties-batch-jobdefinition-firelensconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The image used to start a container. This string is passed directly to the Docker daemon. By
default, images in the Docker Hub registry are available. Other repositories are specified with
either `repository-url/image:tag` or `repository-url/image@digest`. Up to
255 letters (uppercase and lowercase), numbers, hyphens, underscores, colons, periods, forward
slashes, and number signs are allowed. This parameter maps to `Image` in the [Create a\
container](https://docs.docker.com/engine/api/v1.35) section of the [Docker\
Remote API](https://docs.docker.com/engine/api/v1.35) and the `IMAGE` parameter of the [_docker_\
_run_](https://docs.docker.com/engine/reference/run).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinuxParameters`

Linux-specific modifications that are applied to the container, such as Linux kernel
capabilities. For more information, see [KernelCapabilities](../../../../reference/amazonecs/latest/apireference/api-kernelcapabilities.md).

_Required_: No

_Type_: [LinuxParameters](aws-properties-batch-jobdefinition-linuxparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The log configuration specification for the container.

This parameter maps to `LogConfig` in the [Create a\
container](https://docs.docker.com/engine/api/v1.35) section of the [Docker\
Remote API](https://docs.docker.com/engine/api/v1.35) and the `--log-driver` option to [docker\
run](https://docs.docker.com/engine/reference/run).

By default, containers use the same logging driver that the Docker daemon uses. However the
container can use a different logging driver than the Docker daemon by specifying a log driver
with this parameter in the container definition. To use a different logging driver for a
container, the log system must be configured properly on the container instance (or on a
different log server for remote logging options). For more information about the options for
different supported log drivers, see [Configure logging drivers](https://docs.docker.com/engine/admin/logging/overview)
in the _Docker documentation_.

###### Note

Amazon ECS currently supports a subset of the logging drivers available to the Docker daemon
(shown in the `LogConfiguration` data type). Additional log drivers may be available
in future releases of the Amazon ECS container agent.

This parameter requires version 1.18 of the Docker Remote API or greater on your container
instance. To check the Docker Remote API version on your container instance, log in to your
container instance and run the following command: sudo docker version `--format
    '{{.Server.APIVersion}}'`

###### Note

The Amazon ECS container agent running on a container instance must register the logging drivers
available on that instance with the `ECS_AVAILABLE_LOGGING_DRIVERS` environment
variable before containers placed on that instance can use these log configuration options. For
more information, see [Amazon ECS container agent\
configuration](../../../amazonecs/latest/developerguide/ecs-agent-config.md) in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: [LogConfiguration](aws-properties-batch-jobdefinition-logconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountPoints`

The mount points for data volumes in your container.

This parameter maps to `Volumes` in the [Create a\
container](https://docs.docker.com/engine/api/v1.35) section of the [Docker\
Remote API](https://docs.docker.com/engine/api/v1.35) and the --volume option to [docker\
run](https://docs.docker.com/engine/reference/run).

Windows containers can mount whole directories on the same drive as
`$env:ProgramData`. Windows containers can't mount directories on a different drive,
and mount point can't be across drives.

_Required_: No

_Type_: Array of [MountPoint](aws-properties-batch-jobdefinition-mountpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a container. The name can be used as a unique identifier to target your
`dependsOn` and `Overrides` objects.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Privileged`

When this parameter is `true`, the container is given elevated privileges on the
host container instance (similar to the `root` user). This parameter maps to
`Privileged` in the [Create a\
container](https://docs.docker.com/engine/api/v1.35) section of the [Docker\
Remote API](https://docs.docker.com/engine/api/v1.35) and the `--privileged` option to [docker\
run](https://docs.docker.com/engine/reference/run).

###### Note

This parameter is not supported for Windows containers or tasks run on Fargate.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadonlyRootFilesystem`

When this parameter is true, the container is given read-only access to its root file
system. This parameter maps to `ReadonlyRootfs` in the [Create a\
container](https://docs.docker.com/engine/api/v1.35) section of the [Docker\
Remote API](https://docs.docker.com/engine/api/v1.35) and the `--read-only` option to [docker\
run](https://docs.docker.com/engine/reference/run).

###### Note

This parameter is not supported for Windows containers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryCredentials`

The private repository authentication credentials to use.

_Required_: No

_Type_: [RepositoryCredentials](aws-properties-batch-jobdefinition-repositorycredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRequirements`

The type and amount of a resource to assign to a container. The only supported resource is a
GPU.

_Required_: No

_Type_: Array of [ResourceRequirement](aws-properties-batch-jobdefinition-resourcerequirement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Secrets`

The secrets to pass to the container. For more information, see [Specifying Sensitive\
Data](../../../amazonecs/latest/developerguide/specifying-sensitive-data.md) in the Amazon Elastic Container Service Developer Guide.

_Required_: No

_Type_: Array of [Secret](aws-properties-batch-jobdefinition-secret.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ulimits`

A list of `ulimits` to set in the container. If a `ulimit` value is
specified in a task definition, it overrides the default values set by Docker. This parameter
maps to `Ulimits` in the [Create a\
container](https://docs.docker.com/engine/api/v1.35) section of the [Docker\
Remote API](https://docs.docker.com/engine/api/v1.35) and the `--ulimit` option to [docker\
run](https://docs.docker.com/engine/reference/run).

Amazon ECS tasks hosted on Fargate use the default resource limit values set by the operating
system with the exception of the nofile resource limit parameter which Fargate overrides. The
`nofile` resource limit sets a restriction on the number of open files that a
container can use. The default `nofile` soft limit is `1024` and the
default hard limit is `65535`.

This parameter requires version 1.18 of the Docker Remote API or greater on your container
instance. To check the Docker Remote API version on your container instance, log in to your
container instance and run the following command: sudo docker version `--format
    '{{.Server.APIVersion}}'`

###### Note

This parameter is not supported for Windows containers.

_Required_: No

_Type_: Array of [Ulimit](aws-properties-batch-jobdefinition-ulimit.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

The user to use inside the container. This parameter maps to User in the Create a container
section of the Docker Remote API and the --user option to docker run.

###### Note

When running tasks using the `host` network mode, don't run containers using the
`root user (UID 0)`. We recommend using a non-root user for better security.

You can specify the `user` using the following formats. If specifying a UID or
GID, you must specify it as a positive integer.

- `user`

- `user:group`

- `uid`

- `uid:gid`

- `user:gi`

- `uid:group`

###### Note

This parameter is not supported for Windows containers.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TaskContainerDependency

Tmpfs

All content copied from https://docs.aws.amazon.com/.
