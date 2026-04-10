This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ExcludePeriodConfiguration

The exclude period of `TimeRangeFilter` or `RelativeDatesFilter`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amount" : Number,
  "Granularity" : String,
  "Status" : String
}

```

### YAML

```yaml

  Amount: Number
  Granularity: String
  Status: String

```

## Properties

`Amount`

The amount or number of the exclude period.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Granularity`

The granularity or unit (day, month, year) of the exclude period.

_Required_: Yes

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the exclude period. Choose from the following options:

- `ENABLED`

- `DISABLED`

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Entity

ExplicitHierarchy

All content copied from https://docs.aws.amazon.com/.
