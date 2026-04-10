This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition PortConfiguration

A set of port ranges that can be opened on the container. A process that's running in the
container can bind to a port number, making it accessible to inbound traffic when it's mapped
to a container fleet's connection port.

Each container port range specifies a network protocol. When the configuration supports
more than one protocol, we recommend that you use a different range for each protocol. If your
ranges have overlapping port numbers, Amazon GameLift Servers maps a duplicated container port number to
different connection ports. For example, if you include 1935 in port ranges for both TCP and
UDP, it might result in the following mappings:

- container port 1935 (tcp) => connection port 2001

- container port 1935 (udp) => connection port 2002

**Part of:** [GameServerContainerDefinition](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinition.md),
[GameServerContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinitioninput.md),
[SupportContainerDefinition](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinition.md),
[SupportContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinitioninput.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerPortRanges" : [ ContainerPortRange, ... ]
}

```

### YAML

```yaml

  ContainerPortRanges:
    - ContainerPortRange

```

## Properties

`ContainerPortRanges`

A set of one or more container port number ranges. The ranges can't overlap if the ranges'
network protocols are the same. Overlapping ranges with different protocols is allowed but not
recommended.

_Required_: Yes

_Type_: Array of [ContainerPortRange](aws-properties-gamelift-containergroupdefinition-containerportrange.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GameServerContainerDefinition

SupportContainerDefinition

All content copied from https://docs.aws.amazon.com/.
