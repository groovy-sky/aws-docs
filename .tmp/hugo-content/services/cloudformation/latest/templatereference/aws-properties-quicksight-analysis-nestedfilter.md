This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis NestedFilter

A `NestedFilter` filters data with a subset of data that is defined by the nested inner filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "FilterId" : String,
  "IncludeInnerSet" : Boolean,
  "InnerFilter" : InnerFilter
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  FilterId: String
  IncludeInnerSet: Boolean
  InnerFilter:
    InnerFilter

```

## Properties

`Column`

The column that the filter is applied to.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterId`

An identifier that uniquely identifies a filter within a dashboard, analysis, or template.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeInnerSet`

A boolean condition to include or exclude the subset that is defined by the values of the nested inner filter.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InnerFilter`

The `InnerFilter` defines the subset of data to be used with the `NestedFilter`.

_Required_: Yes

_Type_: [InnerFilter](aws-properties-quicksight-analysis-innerfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NegativeValueConfiguration

NullValueFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
