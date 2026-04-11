This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::QueueLimitAssociation

Associates a limit with a particular queue. After the limit is associated, all workers
for jobs that specify the limit associated with the queue are subject to the limit. You
can't associate two limits with the same `amountRequirementName` to the same
queue.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::QueueLimitAssociation",
  "Properties" : {
      "FarmId" : String,
      "LimitId" : String,
      "QueueId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::QueueLimitAssociation
Properties:
  FarmId: String
  LimitId: String
  QueueId: String

```

## Properties

`FarmId`

The unique identifier of the farm that contains the queue-limit association.

_Required_: Yes

_Type_: String

_Pattern_: `^farm-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LimitId`

The unique identifier of the limit in the association.

_Required_: Yes

_Type_: String

_Pattern_: `^limit-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueueId`

The unique identifier of the queue in the association.

_Required_: Yes

_Type_: String

_Pattern_: `^queue-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier for the queue-limit association.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Deadline::QueueFleetAssociation

AWS::Deadline::StorageProfile

All content copied from https://docs.aws.amazon.com/.
