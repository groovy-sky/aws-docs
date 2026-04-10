This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition ContainerEnvironment

An environment variable to set inside a container, in the form of a key-value pair.

**Part of:** [GameServerContainerDefinition](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinition.md),
[GameServerContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinitioninput.md),
[SupportContainerDefinition](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinition.md),
[SupportContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinitioninput.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The environment variable name.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The environment variable value.

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerDependency

ContainerHealthCheck

All content copied from https://docs.aws.amazon.com/.
