This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe DimensionMapping

Maps source data to a dimension in the target Timestream for LiveAnalytics
table.

For more information, see [Amazon Timestream for LiveAnalytics concepts](../../../timestream/latest/developerguide/concepts.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionName" : String,
  "DimensionValue" : String,
  "DimensionValueType" : String
}

```

### YAML

```yaml

  DimensionName: String
  DimensionValue: String
  DimensionValueType: String

```

## Properties

`DimensionName`

The metadata attributes of the time series. For example, the name and Availability Zone
of an Amazon EC2 instance or the name of the manufacturer of a wind turbine are
dimensions.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionValue`

Dynamic path to the dimension value in the source event.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DimensionValueType`

The data type of the dimension for the time-series data.

_Required_: Yes

_Type_: String

_Allowed values_: `VARCHAR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeadLetterConfig

EcsContainerOverride

All content copied from https://docs.aws.amazon.com/.
