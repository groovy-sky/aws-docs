This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CustomFilterConfiguration

A custom filter that filters based on a single value. This filter can be partially matched.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryValue" : String,
  "MatchOperator" : String,
  "NullOption" : String,
  "ParameterName" : String,
  "SelectAllOptions" : String
}

```

### YAML

```yaml

  CategoryValue: String
  MatchOperator: String
  NullOption: String
  ParameterName: String
  SelectAllOptions: String

```

## Properties

`CategoryValue`

The category value for the filter.

This field is mutually exclusive to `ParameterName`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchOperator`

The match operator that is used to determine if a filter should be applied.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | DOES_NOT_EQUAL | CONTAINS | DOES_NOT_CONTAIN | STARTS_WITH | ENDS_WITH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NullOption`

This option determines how null values should be treated when filtering data.

- `ALL_VALUES`: Include null values in filtered results.

- `NULLS_ONLY`: Only include null values in filtered results.

- `NON_NULLS_ONLY`: Exclude null values from filtered results.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL_VALUES | NULLS_ONLY | NON_NULLS_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterName`

The parameter whose value should be used for the filter value.

This field is mutually exclusive to `CategoryValue`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectAllOptions`

Select all of the values. Null is not the assigned value of select all.

- `FILTER_ALL_VALUES`

_Required_: No

_Type_: String

_Allowed values_: `FILTER_ALL_VALUES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomContentVisual

CustomFilterListConfiguration

All content copied from https://docs.aws.amazon.com/.
