This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard AxisDataOptions

The data options for an axis.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateAxisOptions" : DateAxisOptions,
  "NumericAxisOptions" : NumericAxisOptions
}

```

### YAML

```yaml

  DateAxisOptions:
    DateAxisOptions
  NumericAxisOptions:
    NumericAxisOptions

```

## Properties

`DateAxisOptions`

The options for an axis with a date field.

_Required_: No

_Type_: [DateAxisOptions](aws-properties-quicksight-dashboard-dateaxisoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericAxisOptions`

The options for an axis with a numeric field.

_Required_: No

_Type_: [NumericAxisOptions](aws-properties-quicksight-dashboard-numericaxisoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeAggregationFunction

AxisDisplayMinMaxRange

All content copied from https://docs.aws.amazon.com/.
