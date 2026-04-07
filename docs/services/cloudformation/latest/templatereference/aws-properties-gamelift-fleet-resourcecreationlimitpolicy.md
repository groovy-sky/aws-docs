This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet ResourceCreationLimitPolicy

A policy that limits the number of game sessions a player can create on the same fleet.
This optional policy gives game owners control over how players can consume available game
server resources. A resource creation policy makes the following statement: "An individual
player can create a maximum number of new game sessions within a specified time
period".

The policy is evaluated when a player tries to create a new game session. For example,
assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving
a `CreateGameSession` request, Amazon GameLift checks that the player (identified
by `CreatorId`) has created fewer than 10 game sessions in the past 60
minutes.

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

The policy is evaluated when a player tries to create a new game session. On receiving
a `CreateGameSession` request, Amazon GameLift Servers checks that the player (identified by
`CreatorId`) has created fewer than game session limit in the specified
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

## See also

- [Create GameLift resources using Amazon CloudFront](https://docs.aws.amazon.com/gamelift/latest/developerguide/resources-cloudformation.html) in the _Amazon_
_GameLift Developer Guide_

- [ResourceCreationLimitPolicy](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ResourceCreationLimitPolicy.html) in the _Amazon GameLift API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PlayerGatewayConfiguration

RuntimeConfiguration
