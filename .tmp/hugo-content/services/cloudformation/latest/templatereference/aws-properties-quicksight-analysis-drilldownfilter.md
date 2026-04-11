This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DrillDownFilter

The drill down filter for the column hierarchies.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryFilter" : CategoryDrillDownFilter,
  "NumericEqualityFilter" : NumericEqualityDrillDownFilter,
  "TimeRangeFilter" : TimeRangeDrillDownFilter
}

```

### YAML

```yaml

  CategoryFilter:
    CategoryDrillDownFilter
  NumericEqualityFilter:
    NumericEqualityDrillDownFilter
  TimeRangeFilter:
    TimeRangeDrillDownFilter

```

## Properties

`CategoryFilter`

The category type drill down filter. This filter is used for string type columns.

_Required_: No

_Type_: [CategoryDrillDownFilter](aws-properties-quicksight-analysis-categorydrilldownfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericEqualityFilter`

The numeric equality type drill down filter. This filter is used for number type columns.

_Required_: No

_Type_: [NumericEqualityDrillDownFilter](aws-properties-quicksight-analysis-numericequalitydrilldownfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeRangeFilter`

The time range drill down filter. This filter is used for date time columns.

_Required_: No

_Type_: [TimeRangeDrillDownFilter](aws-properties-quicksight-analysis-timerangedrilldownfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DonutOptions

DropDownControlDisplayOptions

All content copied from https://docs.aws.amazon.com/.
