This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition ContainerPortRange

A set of one or more port numbers that can be opened on the container, and the supported
network protocol.

**Part of:** [ContainerPortConfiguration](../../../../reference/gamelift/latest/apireference/api-containerportconfiguration.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FromPort" : Integer,
  "Protocol" : String,
  "ToPort" : Integer
}

```

### YAML

```yaml

  FromPort: Integer
  Protocol: String
  ToPort: Integer

```

## Properties

`FromPort`

A starting value for the range of allowed port numbers.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The network protocol that these ports support.

_Required_: Yes

_Type_: String

_Allowed values_: `TCP | UDP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToPort`

An ending value for the range of allowed port numbers.
Port numbers are end-inclusive.
This value must be equal to or greater than `FromPort`.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `60000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerMountPoint

GameServerContainerDefinition

All content copied from https://docs.aws.amazon.com/.
