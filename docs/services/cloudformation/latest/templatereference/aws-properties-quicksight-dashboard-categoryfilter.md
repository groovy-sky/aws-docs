This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard CategoryFilter

A `CategoryFilter` filters text values.

For more information, see [Adding text filters](../../../quicksight/latest/user/add-a-text-filter-data-prep.md) in the _Amazon Quick Suite User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "Configuration" : CategoryFilterConfiguration,
  "DefaultFilterControlConfiguration" : DefaultFilterControlConfiguration,
  "FilterId" : String
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  Configuration:
    CategoryFilterConfiguration
  DefaultFilterControlConfiguration:
    DefaultFilterControlConfiguration
  FilterId: String

```

## Properties

`Column`

The column that the filter is applied to.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configuration`

The configuration for a `CategoryFilter`.

_Required_: Yes

_Type_: [CategoryFilterConfiguration](aws-properties-quicksight-dashboard-categoryfilterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultFilterControlConfiguration`

The default configurations for the associated controls. This applies only for filters that are scoped to multiple sheets.

_Required_: No

_Type_: [DefaultFilterControlConfiguration](aws-properties-quicksight-dashboard-defaultfiltercontrolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterId`

An identifier that uniquely identifies a filter within a dashboard, analysis, or template.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CategoryDrillDownFilter

CategoryFilterConfiguration

All content copied from https://docs.aws.amazon.com/.
