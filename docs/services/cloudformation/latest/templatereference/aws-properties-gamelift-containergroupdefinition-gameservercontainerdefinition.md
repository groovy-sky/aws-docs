This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition GameServerContainerDefinition

Describes the game server container in an existing game server container group. A game
server container identifies a container image with your game server build. A game server
container is automatically considered essential; if an essential container fails, the entire
container group restarts.

You can update a container definition and deploy the updates to an existing fleet. When
creating or updating a game server container group definition, use the property
[https://docs.aws.amazon.com/gamelift/latest/apireference/API\_GameServerContainerDefinitionInput](https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameServerContainerDefinitionInput).

**Part of:** [ContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ContainerGroupDefinition.html)

**Returned by:** [DescribeContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeContainerGroupDefinition.html),
[ListContainerGroupDefinitions](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListContainerGroupDefinitions.html),
[UpdateContainerGroupDefinition](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateContainerGroupDefinition.html)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerName" : String,
  "DependsOn" : [ ContainerDependency, ... ],
  "EnvironmentOverride" : [ ContainerEnvironment, ... ],
  "ImageUri" : String,
  "MountPoints" : [ ContainerMountPoint, ... ],
  "PortConfiguration" : PortConfiguration,
  "ResolvedImageDigest" : String,
  "ServerSdkVersion" : String
}

```

### YAML

```yaml

  ContainerName: String
  DependsOn:
    - ContainerDependency
  EnvironmentOverride:
    - ContainerEnvironment
  ImageUri: String
  MountPoints:
    - ContainerMountPoint
  PortConfiguration:
    PortConfiguration
  ResolvedImageDigest: String
  ServerSdkVersion: String

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
container group during startup and shutdown sequences. A container might have dependencies on
multiple containers.

_Required_: No

_Type_: Array of [ContainerDependency](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containergroupdefinition-containerdependency.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentOverride`

A set of environment variables that's passed to the container on startup. See the [ContainerDefinition::environment](../../../../reference/amazonecs/latest/apireference/api-containerdefinition.md#ECS-Type-ContainerDefinition-environment) parameter in the _Amazon Elastic Container Service API_
_Reference_.

_Required_: No

_Type_: Array of [ContainerEnvironment](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containergroupdefinition-containerenvironment.html)

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageUri`

The URI to the image that Amazon GameLift Servers uses when deploying this container to a container fleet.
For a more specific identifier, see `ResolvedImageDigest`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_\.@\/:]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountPoints`

A mount point that binds a path inside the container to a file or directory on the host
system and lets it access the file or directory.

_Required_: No

_Type_: Array of [ContainerMountPoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containergroupdefinition-containermountpoint.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortConfiguration`

The set of ports that are available to bind to processes in the container. For example, a
game server process requires a container port to allow game clients to connect to it.
Container ports aren't directly accessed by inbound traffic. Amazon GameLift Servers maps these container
ports to externally accessible connection ports, which are assigned as needed from the
container fleet's `ConnectionPortRange`.

_Required_: No

_Type_: [PortConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-containergroupdefinition-portconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolvedImageDigest`

A unique and immutable identifier for the container image. The digest is a SHA 256 hash of
the container image manifest.

_Required_: No

_Type_: String

_Pattern_: `^sha256:[a-fA-F0-9]{64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSdkVersion`

The Amazon GameLift Servers server SDK version that the game server is integrated with. Only game servers
using 5.2.0 or higher are compatible with container fleets.

_Required_: Yes

_Type_: String

_Pattern_: `^\d+\.\d+\.\d+$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ContainerPortRange

PortConfiguration
