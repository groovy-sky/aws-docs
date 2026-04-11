This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TimeRangeFilterValue

The value of a time range filter.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameter" : String,
  "RollingDate" : RollingDateConfiguration,
  "StaticValue" : String
}

```

### YAML

```yaml

  Parameter: String
  RollingDate:
    RollingDateConfiguration
  StaticValue: String

```

## Properties

`Parameter`

The parameter type input value.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollingDate`

The rolling date input value.

_Required_: No

_Type_: [RollingDateConfiguration](aws-properties-quicksight-analysis-rollingdateconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticValue`

The static input value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeRangeFilter

TooltipItem

All content copied from https://docs.aws.amazon.com/.
