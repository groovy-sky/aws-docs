This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition SupportContainerDefinition

Describes a support container in a container group. A support container might be in a game
server container group or a per-instance container group. Support containers don't run game
server processes.

You can update a support container definition and deploy the updates to an existing fleet.
When creating or updating a game server container group definition, use the property
[GameServerContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinitioninput.md).

**Part of:** [ContainerGroupDefinition](../../../../reference/gamelift/latest/apireference/api-containergroupdefinition.md)

**Returned by:** [DescribeContainerGroupDefinition](../../../../reference/gamelift/latest/apireference/api-describecontainergroupdefinition.md),
[ListContainerGroupDefinitions](../../../../reference/gamelift/latest/apireference/api-listcontainergroupdefinitions.md),
[UpdateContainerGroupDefinition](../../../../reference/gamelift/latest/apireference/api-updatecontainergroupdefinition.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerName" : String,
  "DependsOn" : [ ContainerDependency, ... ],
  "EnvironmentOverride" : [ ContainerEnvironment, ... ],
  "Essential" : Boolean,
  "HealthCheck" : ContainerHealthCheck,
  "ImageUri" : String,
  "MemoryHardLimitMebibytes" : Integer,
  "MountPoints" : [ ContainerMountPoint, ... ],
  "PortConfiguration" : PortConfiguration,
  "ResolvedImageDigest" : String,
  "Vcpu" : Number
}

```

### YAML

```yaml

  ContainerName: String
  DependsOn:
    - ContainerDependency
  EnvironmentOverride:
    - ContainerEnvironment
  Essential: Boolean
  HealthCheck:
    ContainerHealthCheck
  ImageUri: String
  MemoryHardLimitMebibytes: Integer
  MountPoints:
    - ContainerMountPoint
  PortConfiguration:
    PortConfiguration
  ResolvedImageDigest: String
  Vcpu: Number

```

## Properties

`ContainerName`

The container definition identifier. Container names are unique within a container group
definition.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DependsOn`

Indicates that the container relies on the status of other containers in the same
container group during its startup and shutdown sequences. A container might have dependencies
on multiple containers.

_Required_: No

_Type_: Array of [ContainerDependency](aws-properties-gamelift-containergroupdefinition-containerdependency.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentOverride`

A set of environment variables that's passed to the container on startup. See the [ContainerDefinition::environment](../../../../reference/amazonecs/latest/apireference/api-containerdefinition.md#ECS-Type-ContainerDefinition-environment) parameter in the _Amazon Elastic Container Service API_
_Reference_.

_Required_: No

_Type_: Array of [ContainerEnvironment](aws-properties-gamelift-containergroupdefinition-containerenvironment.md)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Essential`

Indicates whether the container is vital to the container group. If an essential container
fails, the entire container group restarts.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheck`

A configuration for a non-terminal health check. A support container automatically
restarts if it stops functioning or if it fails this health check.

_Required_: No

_Type_: [ContainerHealthCheck](aws-properties-gamelift-containergroupdefinition-containerhealthcheck.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageUri`

The URI to the image that Amazon GameLift Servers deploys to a container fleet. For a more specific
identifier, see `ResolvedImageDigest`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_\.@\/:]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryHardLimitMebibytes`

The amount of memory that Amazon GameLift Servers makes available to the container. If memory limits
aren't set for an individual container, the container shares the container group's total
memory allocation.

**Related data type:** [ContainerGroupDefinition TotalMemoryLimitMebibytes](../../../../reference/gamelift/latest/apireference/api-containergroupdefinition.md)

_Required_: No

_Type_: Integer

_Minimum_: `4`

_Maximum_: `1024000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountPoints`

A mount point that binds a path inside the container to a file or directory on the host
system and lets it access the file or directory.

_Required_: No

_Type_: Array of [ContainerMountPoint](aws-properties-gamelift-containergroupdefinition-containermountpoint.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortConfiguration`

A set of ports that allow access to the container from external users. Processes running
in the container can bind to a one of these ports. Container ports aren't directly accessed by
inbound traffic. Amazon GameLift Servers maps these container ports to externally accessible connection ports,
which are assigned as needed from the container fleet's
`ConnectionPortRange`.

_Required_: No

_Type_: [PortConfiguration](aws-properties-gamelift-containergroupdefinition-portconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolvedImageDigest`

A unique and immutable identifier for the container image. The digest is a SHA 256 hash of
the container image manifest.

_Required_: No

_Type_: String

_Pattern_: `^sha256:[a-fA-F0-9]{64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vcpu`

The number of vCPU units that are reserved for the container. If no resources are
reserved, the container shares the total vCPU limit for the container group.

**Related data type:** [ContainerGroupDefinition TotalVcpuLimit](../../../../reference/gamelift/latest/apireference/api-containergroupdefinition.md)

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
