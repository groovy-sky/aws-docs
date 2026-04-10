This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard AttributeAggregationFunction

Aggregation for attributes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SimpleAttributeAggregation" : String,
  "ValueForMultipleValues" : String
}

```

### YAML

```yaml

  SimpleAttributeAggregation: String
  ValueForMultipleValues: String

```

## Properties

`SimpleAttributeAggregation`

The built-in aggregation functions for attributes.

- `UNIQUE_VALUE`: Returns the unique value for a field, aggregated by the dimension fields.

_Required_: No

_Type_: String

_Allowed values_: `UNIQUE_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueForMultipleValues`

Used by the `UNIQUE_VALUE` aggregation function. If there are multiple values for the field used by the aggregation, the value for this property will be returned instead. Defaults to '\*'.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetOptions

AxisDataOptions

All content copied from https://docs.aws.amazon.com/.
