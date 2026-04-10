This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable ReadProvisionedThroughputSettings

Allows you to specify the read capacity settings for a replica table or a replica global
secondary index when the `BillingMode` is set to `PROVISIONED`. You
must specify a value for either `ReadCapacityUnits` or
`ReadCapacityAutoScalingSettings`, but not both. You can switch between fixed
capacity and auto scaling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReadCapacityAutoScalingSettings" : CapacityAutoScalingSettings,
  "ReadCapacityUnits" : Integer
}

```

### YAML

```yaml

  ReadCapacityAutoScalingSettings:
    CapacityAutoScalingSettings
  ReadCapacityUnits: Integer

```

## Properties

`ReadCapacityAutoScalingSettings`

Specifies auto scaling settings for the replica table or global secondary index.

_Required_: No

_Type_: [CapacityAutoScalingSettings](aws-properties-dynamodb-globaltable-capacityautoscalingsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadCapacityUnits`

Specifies a fixed read capacity for the replica table or global secondary index.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReadOnDemandThroughputSettings

ReplicaGlobalSecondaryIndexSpecification

All content copied from https://docs.aws.amazon.com/.
