This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameSessionQueue GameSessionQueueDestination

A fleet or alias designated in a game session queue. Queues fulfill requests for new
game sessions by placing a new game session on any of the queue's destinations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationArn" : String
}

```

### YAML

```yaml

  DestinationArn: String

```

## Properties

`DestinationArn`

The Amazon Resource Name (ARN) that is assigned to fleet or fleet alias. ARNs, which
include a fleet ID or alias ID and a Region name, provide a unique identifier across all
Regions.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterConfiguration

PlayerLatencyPolicy

All content copied from https://docs.aws.amazon.com/.
