This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameSessionQueue FilterConfiguration

A list of fleet locations where a game session queue can place new game sessions. You
can use a filter to temporarily turn off placements for specific locations. For queues
that have multi-location fleets, you can use a filter configuration allow placement with
some, but not all of these locations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedLocations" : [ String, ... ]
}

```

### YAML

```yaml

  AllowedLocations:
    - String

```

## Properties

`AllowedLocations`

A list of locations to allow game session placement in, in the form of AWS Region
codes such as `us-west-2`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GameLift::GameSessionQueue

GameSessionQueueDestination

All content copied from https://docs.aws.amazon.com/.
