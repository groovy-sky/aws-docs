This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::QuotaShare QuotaShareCapacityLimit

Defines the capacity limit for a quota share, or the type and maximum quantity of a particular resource that can be allocated to jobs in the quota share
without borrowing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityUnit" : String,
  "MaxCapacity" : Integer
}

```

### YAML

```yaml

  CapacityUnit: String
  MaxCapacity: Integer

```

## Properties

`CapacityUnit`

The unit of compute capacity for the capacityLimit. For example, `ml.m5.large`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacity`

The maximum capacity available for the quota share. This value represents the maximum quantity of a resource that can be allocated to jobs in the quota share
without borrowing.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Batch::QuotaShare

QuotaSharePreemptionConfiguration
