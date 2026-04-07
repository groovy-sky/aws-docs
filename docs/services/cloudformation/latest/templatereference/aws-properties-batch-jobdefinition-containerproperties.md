This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition ContainerProperties

Container properties are used for Amazon ECS based job definitions. These properties to describe the
container that's launched as part of a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "EnableExecuteCommand" : Boolean,
  "Environment" : [ Environment, ... ],
  "EphemeralStorage" : EphemeralStorage,
  "ExecutionRoleArn" : String,
  "FargatePlatformConfiguration" : FargatePlatformConfiguration,
  "Image" : String,
  "JobRoleArn" : String,
  "LinuxParameters" : LinuxParameters,
  "LogConfiguration" : LogConfiguration,
  "Memory" : Integer,
  "MountPoints" : [ MountPoint, ... ],
  "NetworkConfiguration" : NetworkConfiguration,
  "Privileged" : Boolean,
  "ReadonlyRootFilesystem" : Boolean,
  "RepositoryCredentials" : RepositoryCredentials,
  "ResourceRequirements" : [ ResourceRequirement, ... ],
  "RuntimePlatform" : RuntimePlatform,
  "Secrets" : [ Secret, ... ],
  "Ulimits" : [ Ulimit, ... ],
  "User" : String,
  "Vcpus" : Integer,
  "Volumes" : [ Volume, ... ]
}

```

### YAML

```yaml

  Command:
    - String
  EnableExecuteCommand: Boolean
  Environment:
    - Environment
  EphemeralStorage:
    EphemeralStorage
  ExecutionRoleArn: String
  FargatePlatformConfiguration:
    FargatePlatformConfiguration
  Image: String
  JobRoleArn: String
  LinuxParameters:
    LinuxParameters
  LogConfiguration:
    LogConfiguration
  Memory: Integer
  MountPoints:
    - MountPoint
  NetworkConfiguration:
    NetworkConfiguration
  Privileged: Boolean
  ReadonlyRootFilesystem: Boolean
  RepositoryCredentials:
    RepositoryCredentials
  ResourceRequirements:
    - ResourceRequirement
  RuntimePlatform:
    RuntimePlatform
  Secrets:
    - Secret
  Ulimits:
    - Ulimit
  User: String
  Vcpus: Integer
  Volumes:
    - Volume

```

## Properties

`Command`

The command that's passed to the container. This parameter maps to `Cmd` in the
[Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the `COMMAND`
parameter to [docker run](https://docs.docker.com/engine/reference/run). For more information, see
[https://docs.docker.com/engine/reference/builder/#cmd](https://docs.docker.com/engine/reference/builder).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableExecuteCommand`

Determines whether execute command functionality is turned on for this task. If `true`, execute
command functionality is turned on all the containers in the task.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Environment`

The environment variables to pass to a container. This parameter maps to `Env` in
the [Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the
`--env` option to [docker run](https://docs.docker.com/engine/reference/run).

###### Important

We don't recommend using plaintext environment variables for sensitive information, such as
credential data.

###### Note

Environment variables cannot start with " `AWS_BATCH`". This naming
convention is reserved for variables that AWS Batch sets.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-environment.html) of [Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-environment.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EphemeralStorage`

The amount of ephemeral storage to allocate for the task. This parameter is used to expand
the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on
AWS Fargate.

_Required_: No

_Type_: [EphemeralStorage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-ephemeralstorage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the execution role that AWS Batch can assume. For jobs that run on Fargate
resources, you must provide an execution role. For more information, see [AWS Batch execution IAM role](https://docs.aws.amazon.com/batch/latest/userguide/execution-IAM-role.html)
in the _AWS Batch User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FargatePlatformConfiguration`

The platform configuration for jobs that are running on Fargate resources. Jobs that are
running on Amazon EC2 resources must not specify this parameter.

_Required_: No

_Type_: [FargatePlatformConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-fargateplatformconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

Required. The image used to start a container. This string is passed directly to the
Docker daemon. Images in the Docker Hub registry are available by default. Other repositories are
specified with
`repository-url/image:tag`.
It can be 255 characters long. It can contain uppercase and lowercase letters, numbers,
hyphens (-), underscores (\_), colons (:), periods (.), forward slashes (/), and number signs (#). This parameter maps to `Image` in the
[Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the `IMAGE`
parameter of [docker run](https://docs.docker.com/engine/reference/run).

###### Note

Docker image architecture must match the processor architecture of the compute resources
that they're scheduled on. For example, ARM-based Docker images can only run on ARM-based
compute resources.

- Images in Amazon ECR Public repositories use the full `registry/repository[:tag]` or
`registry/repository[@digest]` naming conventions. For example,
`public.ecr.aws/registry_alias/my-web-app:latest`.

- Images in Amazon ECR repositories use the full registry and repository URI (for example,
`123456789012.dkr.ecr.<region-name>.amazonaws.com/<repository-name>`).

- Images in official repositories on Docker Hub use a single name (for example,
`ubuntu` or `mongo`).

- Images in other repositories on Docker Hub are qualified with an organization name (for
example, `amazon/amazon-ecs-agent`).

- Images in other online repositories are qualified further by a domain name (for example,
`quay.io/assemblyline/ubuntu`).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobRoleArn`

The Amazon Resource Name (ARN) of the IAM role that the container can assume for AWS permissions. For more
information, see [IAM roles for tasks](../../../amazonecs/latest/developerguide/task-iam-roles.md) in the
_Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinuxParameters`

Linux-specific modifications that are applied to the container, such as details for device
mappings.

_Required_: No

_Type_: [LinuxParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-linuxparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The log configuration specification for the container.

This parameter maps to `LogConfig` in the [Create a container](https://docs.docker.com/engine/api/v1.23)
section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the `--log-driver` option to [docker run](https://docs.docker.com/engine/reference/run). By default, containers use the same logging
driver that the Docker daemon uses. However the container might use a different logging driver
than the Docker daemon by specifying a log driver with this parameter in the container
definition. To use a different logging driver for a container, the log system must be configured
properly on the container instance (or on a different log server for remote logging options). For
more information on the options for different supported log drivers, see [Configure logging drivers](https://docs.docker.com/engine/admin/logging/overview)
in the Docker documentation.

###### Note

AWS Batch currently supports a subset of the logging drivers available to the Docker daemon
(shown in the [LogConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties-logconfiguration.html) data type).

This parameter requires version 1.18 of the Docker Remote API or greater on your
container instance. To check the Docker Remote API version on your container instance, log in to your
container instance and run the following command: `sudo docker version | grep "Server API version"`

###### Note

The Amazon ECS container agent running on a container instance must register the logging drivers
available on that instance with the `ECS_AVAILABLE_LOGGING_DRIVERS` environment
variable before containers placed on that instance can use these log configuration options. For
more information, see [Amazon ECS container agent\
configuration](../../../amazonecs/latest/developerguide/ecs-agent-config.md) in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: [LogConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-logconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Memory`

This parameter is deprecated, use `resourceRequirements` to specify the memory
requirements for the job definition. It's not supported for jobs running on Fargate resources.
For jobs that run on Amazon EC2 resources, it specifies the memory hard limit (in MiB) for a
container. If your container attempts to exceed the specified number, it's terminated. You must
specify at least 4 MiB of memory for a job using this parameter. The memory hard limit can be
specified in several places. It must be specified for each node at least once.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountPoints`

The mount points for data volumes in your container. This parameter maps to
`Volumes` in the [Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23)
and the `--volume` option to [docker\
run](https://docs.docker.com/engine/reference/run).

_Required_: No

_Type_: Array of [MountPoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-mountpoint.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

The network configuration for jobs that are running on Fargate resources. Jobs that are
running on Amazon EC2 resources must not specify this parameter.

_Required_: No

_Type_: [NetworkConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-networkconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Privileged`

When this parameter is true, the container is given elevated permissions on the host
container instance (similar to the `root` user). This parameter maps to
`Privileged` in the [Create a container](https://docs.docker.com/engine/api/v1.23) section of the
[Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the `--privileged` option to [docker run](https://docs.docker.com/engine/reference/run). The default value is false.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources and
shouldn't be provided, or specified as false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadonlyRootFilesystem`

When this parameter is true, the container is given read-only access to its root file
system. This parameter maps to `ReadonlyRootfs` in the
[Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the
`--read-only` option to `docker run`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryCredentials`

The private repository authentication credentials to use.

_Required_: No

_Type_: [RepositoryCredentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-repositorycredentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRequirements`

The type and amount of resources to assign to a container. The supported resources include
`GPU`, `MEMORY`, and `VCPU`.

_Required_: No

_Type_: Array of [ResourceRequirement](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-resourcerequirement.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimePlatform`

An object that represents the compute environment architecture for AWS Batch jobs on
Fargate.

_Required_: No

_Type_: [RuntimePlatform](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-runtimeplatform.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Secrets`

The secrets for the container. For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html) in the
_AWS Batch User Guide_.

_Required_: No

_Type_: Array of [Secret](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-secret.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ulimits`

A list of `ulimits` to set in the container. This parameter maps to
`Ulimits` in the [Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23)
and the `--ulimit` option to [docker\
run](https://docs.docker.com/engine/reference/run).

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources and
shouldn't be provided.

_Required_: No

_Type_: Array of [Ulimit](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-ulimit.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`User`

The user name to use inside the container. This parameter maps to `User` in the
[Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the `--user`
option to [docker run](https://docs.docker.com/engine/reference/run).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vcpus`

This parameter is deprecated, use `resourceRequirements` to specify the vCPU
requirements for the job definition. It's not supported for jobs running on Fargate resources.
For jobs running on Amazon EC2 resources, it specifies the number of vCPUs reserved for the
job.

Each vCPU is equivalent to 1,024 CPU shares. This parameter maps to `CpuShares`
in the [Create a container](https://docs.docker.com/engine/api/v1.23) section of the [Docker Remote API](https://docs.docker.com/engine/api/v1.23) and the
`--cpu-shares` option to [docker run](https://docs.docker.com/engine/reference/run). The
number of vCPUs must be specified but can be specified in several places. You must specify it at
least once for each node.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Volumes`

A list of data volumes used in a job.

_Required_: No

_Type_: Array of [Volume](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-jobdefinition-volume.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConsumableResourceRequirement

Device
