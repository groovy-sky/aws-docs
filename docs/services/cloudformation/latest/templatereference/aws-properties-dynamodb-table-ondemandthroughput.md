This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table OnDemandThroughput

Sets the maximum number of read and write units for the specified on-demand table. If
you use this property, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxReadRequestUnits" : Integer,
  "MaxWriteRequestUnits" : Integer
}

```

### YAML

```yaml

  MaxReadRequestUnits: Integer
  MaxWriteRequestUnits: Integer

```

## Properties

`MaxReadRequestUnits`

Maximum number of read request units for the specified table.

To specify a maximum `OnDemandThroughput` on your table, set the value of
`MaxReadRequestUnits` as greater than or equal to 1. To remove the
maximum `OnDemandThroughput` that is currently set on your table, set the
value of `MaxReadRequestUnits` to -1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxWriteRequestUnits`

Maximum number of write request units for the specified table.

To specify a maximum `OnDemandThroughput` on your table, set the value of
`MaxWriteRequestUnits` as greater than or equal to 1. To remove the
maximum `OnDemandThroughput` that is currently set on your table, set the
value of `MaxWriteRequestUnits` to -1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocalSecondaryIndex

PointInTimeRecoverySpecification

All content copied from https://docs.aws.amazon.com/.
