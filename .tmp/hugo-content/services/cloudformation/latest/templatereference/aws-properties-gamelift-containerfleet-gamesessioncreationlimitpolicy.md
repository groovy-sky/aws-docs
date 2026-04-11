This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet GameSessionCreationLimitPolicy

A policy that puts limits on the number of game sessions that a player can create
within a specified span of time. With this policy, you can control players' ability to
consume available resources.

The policy is evaluated when a player tries to create a new game session. On receiving
a `CreateGameSession` request, Amazon GameLift Servers checks that the player (identified by
`CreatorId`) has created fewer than game session limit in the specified
time period.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NewGameSessionsPerCreator" : Integer,
  "PolicyPeriodInMinutes" : Integer
}

```

### YAML

```yaml

  NewGameSessionsPerCreator: Integer
  PolicyPeriodInMinutes: Integer

```

## Properties

`NewGameSessionsPerCreator`

A policy that puts limits on the number of game sessions that a player can create
within a specified span of time. With this policy, you can control players' ability to
consume available resources.

The policy evaluates when a player tries to create a new game session. On receiving a
`CreateGameSession` request, Amazon GameLift Servers checks that the player (identified
by `CreatorId`) has created fewer than game session limit in the specified
time period.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyPeriodInMinutes`

The time span used in evaluating the resource creation limit policy.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentDetails

IpPermission

All content copied from https://docs.aws.amazon.com/.
