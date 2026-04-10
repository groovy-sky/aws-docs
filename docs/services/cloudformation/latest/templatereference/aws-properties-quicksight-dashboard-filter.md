This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard Filter

With a `Filter`, you can remove portions of data from a particular visual or view.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryFilter" : CategoryFilter,
  "NestedFilter" : NestedFilter,
  "NumericEqualityFilter" : NumericEqualityFilter,
  "NumericRangeFilter" : NumericRangeFilter,
  "RelativeDatesFilter" : RelativeDatesFilter,
  "TimeEqualityFilter" : TimeEqualityFilter,
  "TimeRangeFilter" : TimeRangeFilter,
  "TopBottomFilter" : TopBottomFilter
}

```

### YAML

```yaml

  CategoryFilter:
    CategoryFilter
  NestedFilter:
    NestedFilter
  NumericEqualityFilter:
    NumericEqualityFilter
  NumericRangeFilter:
    NumericRangeFilter
  RelativeDatesFilter:
    RelativeDatesFilter
  TimeEqualityFilter:
    TimeEqualityFilter
  TimeRangeFilter:
    TimeRangeFilter
  TopBottomFilter:
    TopBottomFilter

```

## Properties

`CategoryFilter`

A `CategoryFilter` filters text values.

For more information, see [Adding text filters](../../../quicksight/latest/user/add-a-text-filter-data-prep.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [CategoryFilter](aws-properties-quicksight-dashboard-categoryfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NestedFilter`

A `NestedFilter` filters data with a subset of data that is defined by the nested inner filter.

_Required_: No

_Type_: [NestedFilter](aws-properties-quicksight-dashboard-nestedfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericEqualityFilter`

A `NumericEqualityFilter` filters numeric values that equal or do not equal a given numeric value.

_Required_: No

_Type_: [NumericEqualityFilter](aws-properties-quicksight-dashboard-numericequalityfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericRangeFilter`

A `NumericRangeFilter` filters numeric values that are either inside or outside a given numeric range.

_Required_: No

_Type_: [NumericRangeFilter](aws-properties-quicksight-dashboard-numericrangefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelativeDatesFilter`

A `RelativeDatesFilter` filters date values that are relative to a given date.

_Required_: No

_Type_: [RelativeDatesFilter](aws-properties-quicksight-dashboard-relativedatesfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeEqualityFilter`

A `TimeEqualityFilter` filters date-time values that equal or do not equal
a given date/time value.

_Required_: No

_Type_: [TimeEqualityFilter](aws-properties-quicksight-dashboard-timeequalityfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeRangeFilter`

A `TimeRangeFilter` filters date-time values that are either inside or outside a given date/time range.

_Required_: No

_Type_: [TimeRangeFilter](aws-properties-quicksight-dashboard-timerangefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopBottomFilter`

A `TopBottomFilter` filters data to the top or bottom values for a given column.

_Required_: No

_Type_: [TopBottomFilter](aws-properties-quicksight-dashboard-topbottomfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilledMapVisual

FilterControl

All content copied from https://docs.aws.amazon.com/.
