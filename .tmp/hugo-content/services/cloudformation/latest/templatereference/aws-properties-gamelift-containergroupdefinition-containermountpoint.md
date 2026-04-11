This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition ContainerMountPoint

A mount point that binds a container to a file or directory on the host system.

**Part of:** [GameServerContainerDefinition](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinition.md), [https://docs.aws.amazon.com/gamelift/latest/apireference/API\_GameServerContainerDefinitionInput.html](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinitioninput.md), [SupportContainerDefinition](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinition.md), [https://docs.aws.amazon.com/gamelift/latest/apireference/API\_SupportContainerDefinitionInput.html](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinitioninput.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessLevel" : String,
  "ContainerPath" : String,
  "InstancePath" : String
}

```

### YAML

```yaml

  AccessLevel: String
  ContainerPath: String
  InstancePath: String

```

## Properties

`AccessLevel`

The type of access for the container.

_Required_: No

_Type_: String

_Allowed values_: `READ_ONLY | READ_AND_WRITE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerPath`

The mount path on the container. If this property isn't set, the instance path is
used.

_Required_: No

_Type_: String

_Pattern_: `^(\/+[^\/]+\/*)+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstancePath`

The path to the source file or directory.

_Required_: Yes

_Type_: String

_Pattern_: `^\/[\s\S]*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerHealthCheck

ContainerPortRange

All content copied from https://docs.aws.amazon.com/.
