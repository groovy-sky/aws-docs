This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::QueueFleetAssociation

Creates an association between a queue and a fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::QueueFleetAssociation",
  "Properties" : {
      "FarmId" : String,
      "FleetId" : String,
      "QueueId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::QueueFleetAssociation
Properties:
  FarmId: String
  FleetId: String
  QueueId: String

```

## Properties

`FarmId`

The identifier of the farm that contains the queue and the fleet.

_Required_: Yes

_Type_: String

_Pattern_: `^farm-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FleetId`

The fleet ID.

_Required_: Yes

_Type_: String

_Pattern_: `^fleet-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueueId`

The queue ID.

_Required_: Yes

_Type_: String

_Pattern_: `^queue-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the queue fleet
associations.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Deadline::QueueEnvironment

AWS::Deadline::QueueLimitAssociation

All content copied from https://docs.aws.amazon.com/.
