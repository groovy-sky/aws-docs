This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable WriteProvisionedThroughputSettings

Specifies an auto scaling policy for write capacity. This policy will be applied to all
replicas. This setting must be specified if `BillingMode` is set to
`PROVISIONED`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WriteCapacityAutoScalingSettings" : CapacityAutoScalingSettings
}

```

### YAML

```yaml

  WriteCapacityAutoScalingSettings:
    CapacityAutoScalingSettings

```

## Properties

`WriteCapacityAutoScalingSettings`

Specifies auto scaling settings for the replica table or global secondary index.

_Required_: No

_Type_: [CapacityAutoScalingSettings](aws-properties-dynamodb-globaltable-capacityautoscalingsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WriteOnDemandThroughputSettings

AWS::DynamoDB::Table

All content copied from https://docs.aws.amazon.com/.
