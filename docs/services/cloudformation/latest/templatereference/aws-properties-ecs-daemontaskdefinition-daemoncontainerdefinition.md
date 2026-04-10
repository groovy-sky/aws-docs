This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::DaemonTaskDefinition DaemonContainerDefinition

A container definition for a daemon task. Daemon container definitions describe the
containers that run as part of a daemon task on container instances managed by capacity
providers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "Cpu" : Integer,
  "DependsOn" : [ ContainerDependency, ... ],
  "EntryPoint" : [ String, ... ],
  "Environment" : [ KeyValuePair, ... ],
  "EnvironmentFiles" : [ EnvironmentFile, ... ],
  "Essential" : Boolean,
  "FirelensConfiguration" : FirelensConfiguration,
  "HealthCheck" : HealthCheck,
  "Image" : String,
  "Interactive" : Boolean,
  "LinuxParameters" : LinuxParameters,
  "LogConfiguration" : LogConfiguration,
  "Memory" : Integer,
  "MemoryReservation" : Integer,
  "MountPoints" : [ MountPoint, ... ],
  "Name" : String,
  "Privileged" : Boolean,
  "PseudoTerminal" : Boolean,
  "ReadonlyRootFilesystem" : Boolean,
  "RepositoryCredentials" : RepositoryCredentials,
  "RestartPolicy" : RestartPolicy,
  "Secrets" : [ Secret, ... ],
  "StartTimeout" : Integer,
  "StopTimeout" : Integer,
  "SystemControls" : [ SystemControl, ... ],
  "Ulimits" : [ Ulimit, ... ],
  "User" : String,
  "WorkingDirectory" : String
}

```

### YAML

```yaml

  Command:
    - String
  Cpu: Integer
  DependsOn:
    - ContainerDependency
  EntryPoint:
    - String
  Environment:
    - KeyValuePair
  EnvironmentFiles:
    - EnvironmentFile
  Essential: Boolean
  FirelensConfiguration:
    FirelensConfiguration
  HealthCheck:
    HealthCheck
  Image: String
  Interactive: Boolean
  LinuxParameters:
    LinuxParameters
  LogConfiguration:
    LogConfiguration
  Memory: Integer
  MemoryReservation: Integer
  MountPoints:
    - MountPoint
  Name: String
  Privileged: Boolean
  PseudoTerminal: Boolean
  ReadonlyRootFilesystem: Boolean
  RepositoryCredentials:
    RepositoryCredentials
  RestartPolicy:
    RestartPolicy
  Secrets:
    - Secret
  StartTimeout: Integer
  StopTimeout: Integer
  SystemControls:
    - SystemControl
  Ulimits:
    - Ulimit
  User: String
  WorkingDirectory: String

```

## Properties

`Command`

The command that's passed to the container.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Cpu`

The number of `cpu` units reserved for the container.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DependsOn`

The dependencies defined for container startup and shutdown. A container can contain
multiple dependencies on other containers in a task definition.

_Required_: No

_Type_: Array of [ContainerDependency](aws-properties-ecs-daemontaskdefinition-containerdependency.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntryPoint`

The entry point that's passed to the container.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Environment`

The environment variables to pass to a container.

_Required_: No

_Type_: Array of [KeyValuePair](aws-properties-ecs-daemontaskdefinition-keyvaluepair.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentFiles`

A list of files containing the environment variables to pass to a container.

_Required_: No

_Type_: Array of [EnvironmentFile](aws-properties-ecs-daemontaskdefinition-environmentfile.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Essential`

If the `essential` parameter of a container is marked as `true`,
and that container fails or stops for any reason, all other containers that are part of
the task are stopped.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirelensConfiguration`

The FireLens configuration for the container. This is used to specify and configure a
log router for container logs.

_Required_: No

_Type_: [FirelensConfiguration](aws-properties-ecs-daemontaskdefinition-firelensconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HealthCheck`

The container health check command and associated configuration parameters for the
container.

_Required_: No

_Type_: [HealthCheck](aws-properties-ecs-daemontaskdefinition-healthcheck.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Image`

The image used to start the container. This string is passed directly to the Docker
daemon. Images in the Docker Hub registry are available by default. Other repositories
are specified with either
`repository-url/image:tag`
or
`repository-url/image@digest`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Interactive`

When this parameter is `true`, you can deploy containerized applications
that require `stdin` or a `tty` to be allocated.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LinuxParameters`

Linux-specific modifications that are applied to the container configuration, such as
Linux kernel capabilities.

_Required_: No

_Type_: [LinuxParameters](aws-properties-ecs-daemontaskdefinition-linuxparameters.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogConfiguration`

The log configuration specification for the container.

_Required_: No

_Type_: [LogConfiguration](aws-properties-ecs-daemontaskdefinition-logconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Memory`

The amount (in MiB) of memory to present to the container. If the container attempts
to exceed the memory specified here, the container is killed.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemoryReservation`

The soft limit (in MiB) of memory to reserve for the container.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MountPoints`

The mount points for data volumes in your container.

_Required_: No

_Type_: Array of [MountPoint](aws-properties-ecs-daemontaskdefinition-mountpoint.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the container. Up to 255 letters (uppercase and lowercase), numbers,
underscores, and hyphens are allowed.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Privileged`

When this parameter is true, the container is given elevated privileges on the host
container instance (similar to the `root` user).

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PseudoTerminal`

When this parameter is `true`, a TTY is allocated.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReadonlyRootFilesystem`

When this parameter is true, the container is given read-only access to its root file
system.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryCredentials`

The private repository authentication credentials to use.

_Required_: No

_Type_: [RepositoryCredentials](aws-properties-ecs-daemontaskdefinition-repositorycredentials.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestartPolicy`

The restart policy for the container. When you set up a restart policy, Amazon ECS can
restart the container without needing to replace the task.

_Required_: No

_Type_: [RestartPolicy](aws-properties-ecs-daemontaskdefinition-restartpolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Secrets`

The secrets to pass to the container.

_Required_: No

_Type_: Array of [Secret](aws-properties-ecs-daemontaskdefinition-secret.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartTimeout`

Time duration (in seconds) to wait before giving up on resolving dependencies for a
container.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StopTimeout`

Time duration (in seconds) to wait before the container is forcefully killed if it
doesn't exit normally on its own.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SystemControls`

A list of namespaced kernel parameters to set in the container.

_Required_: No

_Type_: Array of [SystemControl](aws-properties-ecs-daemontaskdefinition-systemcontrol.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ulimits`

A list of `ulimits` to set in the container.

_Required_: No

_Type_: Array of [Ulimit](aws-properties-ecs-daemontaskdefinition-ulimit.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`User`

The user to use inside the container.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkingDirectory`

The working directory to run commands inside the container in.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerDependency

Device

All content copied from https://docs.aws.amazon.com/.
