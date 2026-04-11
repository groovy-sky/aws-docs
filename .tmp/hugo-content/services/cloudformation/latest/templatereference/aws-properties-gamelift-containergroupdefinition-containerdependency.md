This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerGroupDefinition ContainerDependency

A container's dependency on another container in the same container group. The dependency
impacts how the dependent container is able to start or shut down based the status of the
other container.

For example, _ContainerA_ is configured with the following dependency: a
`START` dependency on _ContainerB_. This means that
_ContainerA_ can't start until _ContainerB_ has
started. It also means that _ContainerA_ must shut down before
_ContainerB_.

**Part of:** [GameServerContainerDefinition](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinition.md),
[GameServerContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-gameservercontainerdefinitioninput.md),
[SupportContainerDefinition](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinition.md),
[SupportContainerDefinitionInput](../../../../reference/gamelift/latest/apireference/api-supportcontainerdefinitioninput.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : String,
  "ContainerName" : String
}

```

### YAML

```yaml

  Condition: String
  ContainerName: String

```

## Properties

`Condition`

The condition that the dependency container must reach before the dependent container can
start. Valid conditions include:

- START - The dependency container must have started.

- COMPLETE - The dependency container has run to completion (exits). Use this condition with
nonessential containers, such as those that run a script and then exit. The dependency
container can't be an essential container.

- SUCCESS - The dependency container has run to completion and exited with a zero status. The
dependency container can't be an essential container.

- HEALTHY - The dependency container has passed its Docker health check. Use this condition with
dependency containers that have health checks configured. This condition is confirmed at
container group startup only.

_Required_: Yes

_Type_: String

_Allowed values_: `START | COMPLETE | SUCCESS | HEALTHY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerName`

A descriptive label for the container definition that this container depends on.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GameLift::ContainerGroupDefinition

ContainerEnvironment

All content copied from https://docs.aws.amazon.com/.
