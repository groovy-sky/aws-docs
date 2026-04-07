This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table PointInTimeRecoverySpecification

The settings used to enable point in time recovery.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PointInTimeRecoveryEnabled" : Boolean,
  "RecoveryPeriodInDays" : Integer
}

```

### YAML

```yaml

  PointInTimeRecoveryEnabled: Boolean
  RecoveryPeriodInDays: Integer

```

## Properties

`PointInTimeRecoveryEnabled`

Indicates whether point in time recovery is enabled (true) or disabled (false) on the
table.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryPeriodInDays`

The number of preceding days for which continuous backups are taken and maintained.
Your table data is only recoverable to any point-in-time from within the configured
recovery period. This parameter is optional. If no value is provided, the value will
default to 35.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `35`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

[PointInTimeRecoverySpecification](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PointInTimeRecoverySpecification.html) in the Amazon DynamoDB API Reference.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OnDemandThroughput

Projection
