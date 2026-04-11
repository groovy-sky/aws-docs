This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameSessionQueue PlayerLatencyPolicy

The queue setting that determines the highest latency allowed for individual
players when placing a game session. When a latency policy is in force, a game session cannot
be placed with any fleet in a Region where a player reports latency higher than the cap.
Latency policies are only enforced when the placement request contains player latency
information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumIndividualPlayerLatencyMilliseconds" : Integer,
  "PolicyDurationSeconds" : Integer
}

```

### YAML

```yaml

  MaximumIndividualPlayerLatencyMilliseconds: Integer
  PolicyDurationSeconds: Integer

```

## Properties

`MaximumIndividualPlayerLatencyMilliseconds`

The maximum latency value that is allowed for any player, in milliseconds. All
policies must have a value set for this property.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyDurationSeconds`

The length of time, in seconds, that the policy is enforced while placing a new game
session. A null value for this property means that the policy is enforced until the
queue times out.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Create a queue](../../../gamelift/latest/developerguide/queues-creating.md) in
the _Amazon GameLift Developer Guide_

- [PlayerLatencyPolicy](../../../../reference/gamelift/latest/apireference/api-playerlatencypolicy.md)
in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GameSessionQueueDestination

PriorityConfiguration

All content copied from https://docs.aws.amazon.com/.
